#
/*
 *  link editor
 */

#include <stdio.h>
#include <stdlib.h>

void delexit();
struct symbol *enter();
struct symbol *lookloc();

void put16(v, f)
short v;
FILE *f;
{
	fwrite(&v, sizeof(v), 1, f);
}

#define	SIGINT	2
#define	ARCMAGIC 0177555
#define	FMAGIC	0407
#define	NMAGIC	0410
#define	IMAGIC	0411

#define	EXTERN	040
#define	UNDEF	00
#define	ABS	01
#define	TEXT	02
#define	DATA	03
#define	BSS	04
#define	COMM	05	/* internal use only */

#define	RABS	00
#define	RTEXT	02
#define	RDATA	04
#define	RBSS	06
#define	REXT	010

#define	RELFLG	01
#define	NROUT	256
#define	NSYM	501
#define	NSYMPR	500

#define	RONLY	0400

char	premeof[] = "Premature EOF on %s";

struct	page {
	short	nuser;
	short	bno;
	short	nibuf;
	short	buff[256];
} page[2];

struct	{
	short	nuser;
	short	bno;
} fpage;

struct	stream {
	short	*ptr;
	short	bno;
	short	nibuf;
	short	size;
	struct	page *pno;
};

struct	stream text;
struct	stream reloc;

struct	archdr {
	char	aname[8];
	short	atime[2];
	char	auid, amode;
	short	asize;
} archdr;

struct	filhdr {
	short	fmagic;
	short	tsize;
	short	dsize;
	short	bsize;
	short	ssize;
	short	entry;
	short	pad;
	short	relflg;
} filhdr;

struct	liblist {
	short	off;
	short	bno;
};

struct	liblist	liblist[NROUT];
struct	liblist	*libp = { &liblist[0] };

struct	symbol {
	char	sname[8];
	char	stype;
	char	spad;
	short	svalue;
};

struct	symbol	cursym;
struct	symbol	symtab[NSYM];
struct	symbol	*hshtab[NSYM+2];
struct	symbol	*symp = { symtab };
struct	symbol	**local[NSYMPR];
struct	symbol	*p_etext;
struct	symbol	*p_edata;
struct	symbol	*p_end;

struct	symbol2 {
	int	num;
	struct symbol *sym;
} local2[NSYMPR];

short	xflag;		/* discard local symbols */
short	Xflag;		/* discard locals starting with 'L' */
short	rflag;		/* preserve relocation bits, don't define common */
short	arflag;		/* original copy of rflag */
short	sflag;		/* discard all symbols */
short	nflag;		/* pure procedure */
short	dflag;		/* define common even with rflag */
short	iflag;		/* I/D space separated */

FILE	*infil;
char	*filname;

short	tsize;
short	dsize;
short	bsize;
short	ssize;
short	nsym;

short	torigin;
short	dorigin;
short	borigin;

short	ctrel;
short	cdrel;
short	cbrel;

short	errlev;
short	delarg	= 4;
char	tfname[]	= "/tmp/lxyyyyy";
FILE	*tout;
FILE	*dout;
FILE	*trout;
FILE	*drout;
FILE	*sout;

struct	symbol	**lookup();
struct	symbol	**slookup();

main(argc, argv)
char **argv;
{
	register short c;
	register char *ap, **p;
	struct symbol **hp;

	if ((signal(SIGINT, 1) & 01) == 0)
		signal(SIGINT, delexit);
	if (argc == 1)
		exit(4);
	p = argv + 1;
	for (c = 1; c<argc; c++) {
		filname = 0;
		ap = *p++;
		if (*ap == '-') switch (ap[1]) {

		case 'u':
			if (++c >= argc)
				error(1, "Bad 'use'");
			if (*(hp = slookup(*p++)) == 0) {
				*hp = symp;
				enter();
			}
			continue;

		case 'l':
			break;

		case 'x':
			xflag++;
			continue;

		case 'X':
			Xflag++;
			continue;

		case 'r':
			rflag++;
			arflag++;
			continue;

		case 's':
			sflag++;
			xflag++;
			continue;

		case 'n':
			nflag++;
			continue;

		case 'd':
			dflag++;
			continue;

		case 'i':
			iflag++;
			continue;
		}
		load1arg(ap);
		fclose(infil);
	}
	middle();
	setupout();
	p = argv+1;
	libp = liblist;
	for (c=1; c<argc; c++) {
		ap = *p++;
		if (*ap == '-') switch (ap[1]) {

		case 'u':
			++c;
			++p;
		default:
			continue;

		case 'l':
			break;
		}
		load2arg(ap);
		fclose(infil);
	}
	finishout();
}

load1arg(acp)
char *acp;
{
	register char *cp;
	register short noff, nbno;

	cp = acp;
	if (getfile(cp)==0) {
		load1(0, 0, 0);
		return;
	}
	nbno = 0;
	noff = 1;
	for (;;) {
		dseek(&text, nbno, noff, sizeof archdr);
		if (text.size <= 0) {
			libp->bno = -1;
			libp++;
			return;
		}
		mget(&archdr, sizeof archdr);
		if (load1(1, nbno, noff + (sizeof archdr) / 2)) {
			libp->bno = nbno;
			libp->off = noff;
			libp++;
		}
		noff += (archdr.asize + sizeof archdr)>>1;
		nbno += (noff >> 8) & 0377;
		noff &= 0377;
	}
}

load1(libflg, bno, off)
{
	register struct symbol *sp, **hp, ***cp;
	struct symbol *ssymp;
	short ndef, nloc;

	readhdr(bno, off);
	ctrel = tsize;
	cdrel += dsize;
	cbrel += bsize;
	ndef = 0;
	nloc = sizeof cursym;
	cp = local;
	ssymp = symp;
	if ((filhdr.relflg&RELFLG)==1) {
		error(0, "No relocation bits");
		return(0);
	}
	off += (sizeof filhdr)/2 + filhdr.tsize + filhdr.dsize;
	dseek(&text, bno, off, filhdr.ssize);
	while (text.size > 0) {
		mget(&cursym, sizeof cursym);
		if ((cursym.stype&EXTERN)==0) {
			if (Xflag==0 || cursym.sname[0]!='L')
				nloc += sizeof cursym;
			continue;
		}
		symreloc();
		hp = lookup();
		if ((sp = *hp) == 0) {
			*hp = enter();
			*cp++ = hp;
			continue;
		}
		if (sp->stype != EXTERN+UNDEF)
			continue;
		if (cursym.stype == EXTERN+UNDEF) {
			if (cursym.svalue > sp->svalue)
				sp->svalue = cursym.svalue;
			continue;
		}
		if (sp->svalue != 0 && cursym.stype == EXTERN+TEXT)
			continue;
		ndef++;
		sp->stype = cursym.stype;
		sp->svalue = cursym.svalue;
	}
	if (libflg==0 || ndef) {
		tsize += filhdr.tsize;
		dsize += filhdr.dsize;
		bsize += filhdr.bsize;
		ssize += nloc;
		return(1);
	}
/*
 * No symbols defined by this library member.
 * Rip out the hash table entries and reset the symbol table.
 */
	symp = ssymp;
	while (cp > local)
		**--cp = 0;
	return(0);
}

middle()
{
	register struct symbol *sp;
	register short t, csize;
	short nund, corigin;

	p_etext = *slookup("_etext");
	p_edata = *slookup("_edata");
	p_end = *slookup("_end");
/*
 * If there are any undefined symbols, save the relocation bits.
 */
	if (rflag==0) for (sp=symtab; sp<symp; sp++)
		if (sp->stype==EXTERN+UNDEF && sp->svalue==0
		 && sp!=p_end && sp!=p_edata && sp!=p_etext) {
			rflag++;
			dflag = 0;
			nflag = 0;
			iflag = 0;
			sflag = 0;
			break;
		}
/*
 * Assign common locations.
 */
	csize = 0;
	if (dflag || rflag==0) {
		for (sp=symtab; sp<symp; sp++)
			if (sp->stype==EXTERN+UNDEF && (t=sp->svalue)!=0) {
				t = (t+1) & ~01;
				sp->svalue = csize;
				sp->stype = EXTERN+COMM;
				csize += t;
			}
		if (p_etext && p_etext->stype==EXTERN+UNDEF) {
			p_etext->stype = EXTERN+TEXT;
			p_etext->svalue = tsize;
		}
		if (p_edata && p_edata->stype==EXTERN+UNDEF) {
			p_edata->stype = EXTERN+DATA;
			p_edata->svalue = dsize;
		}
		if (p_end && p_end->stype==EXTERN+UNDEF) {
			p_end->stype = EXTERN+BSS;
			p_end->svalue = bsize;
		}
	}
/*
 * Now set symbols to their final value
 */
	if (nflag || iflag)
		tsize = (tsize + 077) & ~077;
	dorigin = tsize;
	if (nflag)
		dorigin = (tsize+017777) & ~017777;
	if (iflag)
		dorigin = 0;
	corigin = dorigin + dsize;
	borigin = corigin + csize;
	nund = 0;
	for (sp=symtab; sp<symp; sp++) switch (sp->stype) {
	case EXTERN+UNDEF:
		errlev |= 01;
		if (arflag==0 && sp->svalue==0) {
			if (nund==0)
				printf("Undefined:\n");
			nund++;
			printf("%.8s\n", sp->sname);
		}
		continue;

	case EXTERN+ABS:
	default:
		continue;

	case EXTERN+TEXT:
		sp->svalue += torigin;
		continue;

	case EXTERN+DATA:
		sp->svalue += dorigin;
		continue;

	case EXTERN+BSS:
		sp->svalue += borigin;
		continue;

	case EXTERN+COMM:
		sp->stype = EXTERN+BSS;
		sp->svalue += corigin;
		continue;
	}
	if (sflag || xflag)
		ssize = 0;
	bsize += csize;
	nsym = ssize / (sizeof cursym);
}

setupout()
{
	register char *p;
	register short pid;

	if ((tout = fopen("l.out", "wb")) < 0)
		error(1, "Can't create l.out");
	pid = getpid();
	for (p = &tfname[12]; p > &tfname[7];) {
		*--p = (pid&07) + '0';
		pid >>= 3;
	}
	tcreat(dout, 'a');
	if (sflag==0 || xflag==0)
		tcreat(sout, 'b');
	if (rflag) {
		tcreat(trout, 'c');
		tcreat(drout, 'd');
	}
	filhdr.fmagic = FMAGIC;
	if (nflag)
		filhdr.fmagic = NMAGIC;
	if (iflag)
		filhdr.fmagic = IMAGIC;
	filhdr.tsize = tsize;
	filhdr.dsize = dsize;
	filhdr.bsize = bsize;
	filhdr.ssize = sflag? 0: (ssize + (sizeof cursym)*(symp-symtab));
	filhdr.entry = 0;
	filhdr.pad = 0;
	filhdr.relflg = (rflag==0);
	mput(tout, &filhdr, sizeof filhdr);
	return;
}

tcreat(buf, letter)
FILE *buf;
{
#ifdef _WIN32
	extern const char *convtemp(const char *);
	tfname[6] = letter;
	if (!(buf = fopen(convtemp(tfname), "wb")))
		error(1, "Can't create temp");
#else
	tfname[6] = letter;
	if (!(buf = fopen(tfname, "wb")))
		error(1, "Can't create temp");
#endif
}

load2arg(acp)
char *acp;
{
	register char *cp;
	register struct liblist *lp;

	cp = acp;
	if (getfile(cp) == 0) {
		while (*cp)
			cp++;
		while (cp >= acp && *--cp != '/');
		mkfsym(++cp);
		load2(0, 0);
		return;
	}
	for (lp = libp; lp->bno != -1; lp++) {
		dseek(&text, lp->bno, lp->off, sizeof archdr);
		mget(&archdr, sizeof archdr);
		mkfsym(archdr.aname);
		load2(lp->bno, lp->off + (sizeof archdr) / 2);
	}
	libp = ++lp;
}

load2(bno, off)
{
	register struct symbol *sp;
	register short li, symno;

	readhdr(bno, off);
	ctrel = torigin;
	cdrel += dorigin;
	cbrel += borigin;
/*
 * Reread the symbol table, recording the numbering
 * of symbols for fixing external references.
 */
	li = 0;
	symno = -1;
	off += (sizeof filhdr)/2;
	dseek(&text, bno, off+filhdr.tsize+filhdr.dsize, filhdr.ssize);
	while (text.size > 0) {
		symno++;
		mget(&cursym, sizeof cursym);
		symreloc();
		if ((cursym.stype&EXTERN) == 0) {
			if (!sflag&&!xflag&&(!Xflag||cursym.sname[0]!='L'))
				mput(sout, &cursym, sizeof cursym);
			continue;
		}
		if ((sp = *lookup()) == 0)
			error(1, "internal error: symbol not found");
		if (cursym.stype == EXTERN+UNDEF) {
			if (li >= NSYMPR)
				error(1, "Local symbol overflow");
			local2[li].num = symno;
			local2[li].sym = sp;
			++li;
			continue;
		}
		if (cursym.stype!=sp->stype || cursym.svalue!=sp->svalue) {
			printf("%.8s: ", cursym.sname);
			error(0, "Multiply defined");
		}
	}
	dseek(&text, bno, off, filhdr.tsize);
	dseek(&reloc, bno, off+(filhdr.tsize+filhdr.dsize)/2, filhdr.tsize);
	load2td(li, ctrel, tout, trout);
	dseek(&text, bno, off+(filhdr.tsize/2), filhdr.dsize);
	dseek(&reloc, bno, off+filhdr.tsize+(filhdr.dsize/2), filhdr.dsize);
	load2td(li, cdrel, dout, drout);
	torigin += filhdr.tsize;
	dorigin += filhdr.dsize;
	borigin += filhdr.bsize;
}

load2td(li, creloc, b1, b2)
FILE *b1, *b2;
{
	register short r, t;
	register struct symbol *sp;

	for (;;) {
	/*
	 * The pickup code is copied from "get" for speed.
	 */
		if (--text.size <= 0) {
			if (text.size < 0)
				break;
			text.size++;
			t = get(&text);
		} else if (--text.nibuf < 0) {
			text.nibuf++;
			text.size++;
			t = get(&text);
		} else
			t = *text.ptr++;
		if (--reloc.size <= 0) {
			if (reloc.size < 0)
				error(1, "Relocation error");
			reloc.size++;
			r = get(&reloc);
		} else if (--reloc.nibuf < 0) {
			reloc.nibuf++;
			reloc.size++;
			r = get(&reloc);
		} else
			r = *reloc.ptr++;
		switch (r&016) {

		case RTEXT:
			t += ctrel;
			break;

		case RDATA:
			t += cdrel;
			break;

		case RBSS:
			t += cbrel;
			break;

		case REXT:
			sp = lookloc(li, r);
			if (sp->stype==EXTERN+UNDEF) {
				r = (r&01) + ((nsym+(sp-symtab))<<4) + REXT;
				break;
			}
			t += sp->svalue;
			r = (r&01) + ((sp->stype-(EXTERN+ABS))<<1);
			break;
		}
		if (r&01)
			t -= creloc;
		put16(t, b1);
		if (rflag)
			put16(r, b2);
	}
}

finishout()
{
	register short n, *p;

	if (nflag||iflag) {
		n = torigin;
		while (n&077) {
			n += 2;
			put16(0, tout);
			if (rflag)
				put16(0, trout);
		}
	}
	copy(dout, 'a');
	if (rflag) {
		copy(trout, 'c');
		copy(drout, 'd');
	}
	if (sflag==0) {
		if (xflag==0)
			copy(sout, 'b');
		for (p=(short *)symtab; p < (short *)symp;)
			put16(*p++, tout);
	}
	fflush(tout);
	fclose(tout);
	unlink("a.out");
	link("l.out", "a.out");
	delarg = errlev;
	delexit();
}

void
delexit()
{
	register short c;

	unlink("l.out");
	for (c = 'a'; c <= 'd'; c++) {
		tfname[6] = c;
		unlink(tfname);
	}
	if (delarg==0)
		chmod("a.out", 0777);
	exit(delarg);
}

copy(buf, c)
FILE *buf;
{
	FILE *f;
	register short n;
	char b[512];

	fflush(buf);
	fclose(buf);
	tfname[6] = c;
	f = fopen(tfname, "rb");
	while ((n = fread(b, 1, 512, f)) > 0) {
		fwrite(b, 1, n, tout);
	}
	fclose(f);
}

mkfsym(s)
char *s;
{

	if (sflag || xflag)
		return;
	cp8c(s, cursym.sname);
	cursym.stype = 037;
	cursym.svalue = torigin;
	mput(sout, &cursym, sizeof cursym);
}

mget(aloc, an)
short *aloc;
{
	register short *loc, n;
	register short *p;

	n = an;
	n >>= 1;
	loc = aloc;
	if ((text.nibuf -= n) >= 0) {
		if ((text.size -= n) > 0) {
			p = text.ptr;
			do
				*loc++ = *p++;
			while (--n);
			text.ptr = p;
			return;
		} else
			text.size += n;
	}
	text.nibuf += n;
	do {
		*loc++ = get(&text);
	} while (--n);
}

mput(buf, aloc, an)
FILE *buf;
short *aloc;
{
	register short *loc;
	register short n;

	loc = aloc;
	n = an>>1;
	do {
		put16(*loc++, buf);
	} while (--n);
}

dseek(asp, ab, o, s)
struct stream *asp;
{
	register struct stream *sp;
	register struct page *p;
	register short b;
	short n;

	sp = asp;
	b = ab + ((o>>8) & 0377);
	o &= 0377;
	--sp->pno->nuser;
	if ((p = &page[0])->bno!=b && (p = &page[1])->bno!=b)
		if (p->nuser==0 || (p = &page[0])->nuser==0) {
			if (page[0].nuser==0 && page[1].nuser==0)
				if (page[0].bno < page[1].bno)
					p = &page[0];
			p->bno = b;
			fseek(infil, b << 9, SEEK_SET);
			if ((n = fread(p->buff, 1, 512, infil)>>1) < 0)
				n = 0;
			p->nibuf = n;
		} else
			error(1, "No pages");
	++p->nuser;
	sp->bno = b;
	sp->pno = p;
	sp->ptr = p->buff + o;
	if (s != -1)
		sp->size = (s>>1) & 077777;
	if ((sp->nibuf = p->nibuf-o) <= 0)
		sp->size = 0;
}

get(asp)
struct stream *asp;
{
	register struct stream *sp;

	sp = asp;
	if (--sp->nibuf < 0) {
		dseek(sp, sp->bno+1, 0, -1);
		--sp->nibuf;
	}
	if (--sp->size <= 0) {
		if (sp->size < 0)
			error(1, premeof);
		++fpage.nuser;
		--sp->pno->nuser;
	    sp->pno = (struct page *)&fpage;
	}
	return(*sp->ptr++);
}

getfile(acp)
char *acp;
{
	register char *cp;
	register short c;

	cp = acp;
	archdr.aname[0] = '\0';
	if (cp[0]=='-' && cp[1]=='l') {
		if ((c = cp[2]) == '\0')
			c = 'a';
		cp = "/lib/lib?.a";
		cp[8] = c;
	}
	filname = cp;
	if (!(infil = fopen(cp, "rb")))
		error(1, "cannot open");
	page[0].bno = page[1].bno = -1;
	page[0].nuser = page[1].nuser = 0;
	text.pno = reloc.pno = (struct page *)&fpage;
	fpage.nuser = 2;
	dseek(&text, 0, 0, 2);
	if (text.size <= 0)
		error(1, premeof);
	return(get(&text) == ARCMAGIC);
}

struct symbol **lookup()
{
	short i;
	register struct symbol **hp;
	register char *cp, *cp1;

	i = 0;
	for (cp=cursym.sname; cp < &cursym.sname[8];)
		i = (i<<1) + *cp++;
	for (hp = &hshtab[(i&077777)%NSYM+2]; *hp!=0;) {
		cp1 = (*hp)->sname;
		for (cp=cursym.sname; cp < &cursym.sname[8];)
			if (*cp++ != *cp1++)
				goto no;
		break;
	    no:
		if (++hp >= &hshtab[NSYM+2])
			hp = hshtab;
	}
	return(hp);
}

struct symbol **slookup(s)
char *s;
{
	cp8c(s, cursym.sname);
	cursym.stype = EXTERN+UNDEF;
	cursym.svalue = 0;
	return(lookup());
}

struct symbol *
enter()
{
	register struct symbol *sp;
	
	if ((sp=symp) >= &symtab[NSYM])
		error(1, "Symbol table overflow");
	cp8c(cursym.sname, sp->sname);
	sp->stype = cursym.stype;
	sp->svalue = cursym.svalue;
	symp++;
	return(sp);
}

symreloc()
{
	switch (cursym.stype) {

	case TEXT:
	case EXTERN+TEXT:
		cursym.svalue += ctrel;
		return;

	case DATA:
	case EXTERN+DATA:
		cursym.svalue += cdrel;
		return;

	case BSS:
	case EXTERN+BSS:
		cursym.svalue += cbrel;
		return;

	case EXTERN+UNDEF:
		return;
	}
	if (cursym.stype&EXTERN)
		cursym.stype = EXTERN+ABS;
}

error(n, s)
char *s;
{
	if (filname) {
		printf("%s", filname);
		if (archdr.aname[0])
			printf("(%.8s)", archdr.aname);
		printf(": ");
	}
	printf("%s\n", s);
	if (n)
		delexit();
	errlev = 2;
}

struct symbol *
lookloc(li, r)
{
	register short i;
	register short sn;

	sn = (r>>4) & 07777;
	for (i=0; i<li; ++i)
		if (local2[i].num == sn)
			return(local2[i].sym);
	error(1, "Local symbol botch");
}

readhdr(bno, off)
{
	register short st, sd;

	dseek(&text, bno, off, sizeof filhdr);
	mget(&filhdr, sizeof filhdr);
	if (filhdr.fmagic != FMAGIC)
		error(1, "Bad format");
	st = (filhdr.tsize+01) & ~01;
	filhdr.tsize = st;
	cdrel = -st;
	sd = (filhdr.dsize+01) & ~01;
	cbrel = - (st+sd);
	filhdr.bsize = (filhdr.bsize+01) & ~01;
}

cp8c(from, to)
char *from, *to;
{
	register char *f, *t, *te;

	f = from;
	t = to;
	te = t+8;
	while ((*t++ = *f++) && t<te);
	while (t<te)
		*t++ = 0;
}
