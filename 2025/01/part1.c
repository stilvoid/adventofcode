#include <stdio.h>

int main(int argc, char* argv[]) {
	if(argc!=2) {
		fprintf(stderr, "Need INPUT\n");
		return 1;
	}

	char* fn = argv[1];
	FILE* f = fopen(fn, "r");

	int dial = 50;
	char dir;
	int n;

	int answer = 0;

	while(!feof(f)) {
		fscanf(f, "%c%d\n", &dir, &n);

		if(dir == 'L') {
			dial = (dial - n + 100) % 100;
		} else {
			dial = (dial + n) % 100;
		}

		if(dial == 0) {
			answer++;
		}
	}

	fclose(f);

	printf("Answer: %d\n", answer);
}
