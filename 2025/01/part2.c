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

		while(n>=100) {
			answer++;
			n-=100;
		}

		if(dir == 'L') {
			dial -= n;
			if(dial < 0) {
				if(dial+n != 0) {
					answer++;
				}
				dial+=100;
			}
		} else {
			dial += n;
			if(dial>=100) {
				dial-=100;
				if(dial != 0) {
					answer++;
				}
			}
		}

		if(dial == 0) {
			answer++;
		}
	}

	fclose(f);

	printf("Answer: %d\n", answer);
}
