/* $Id: rd_unsig2.c,v 1.5 1994/06/24 11:19:11 ceriel Exp $ */
/*
 * (c) copyright 1987 by the Vrije Universiteit, Amsterdam, The Netherlands.
 * See the copyright notice in the ACK home directory, in the file "Copyright".
 */
#include "obj.h"

unsigned int
rd_unsigned2(fd)
{
	char buf[2];

	rd_bytes(fd, buf, 2L);
	return uget2(buf);
}
