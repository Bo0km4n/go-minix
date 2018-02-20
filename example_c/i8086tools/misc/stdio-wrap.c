#include <stdio.h>
#include <unistd.h>
#include <fcntl.h>

int main(int argc, char *argv[], char *envp[]) {
	char *nargv[2];
	if (argc < 4) {
		printf("usage: %s exec in out\n", argv[0]);
		return 1;
	}
	nargv[0] = argv[1];
	nargv[1] = NULL;
	close(0);
	open(argv[2], 0);
	close(1);
	creat(argv[3], 0666);
	execve(argv[1], nargv, envp);
}
