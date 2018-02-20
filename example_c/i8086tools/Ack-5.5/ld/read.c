/*
 * (c) copyright 1987 by the Vrije Universiteit, Amsterdam, The Netherlands.
 * See the copyright notice in the ACK home directory, in the file "Copyright".
 */
#ifndef lint
static char rcsid[] = "$Id: read.c,v 3.2 1994/06/24 10:35:08 ceriel Exp $";
#endif

int	infile;	/* The current input file. */

rd_fatal()
{
	fatal("read error");
}
