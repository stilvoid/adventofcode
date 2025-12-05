#include <stdio.h>

typedef struct {
    long start;
    long end;
} Range;

Range ranges[2000];
long ids[2000];

int main(int argc, char* argv[]) {
    char line[100];

    int nranges = 0;
    int nids = 0;

    while(fgets(line, 100, stdin)) {
        if(line[0] == '\n') {
            break;
        }

        sscanf(line, "%ld-%ld", &ranges[nranges].start, &ranges[nranges].end);
        nranges++;
    }

    while(fgets(line, 100, stdin)) {
        sscanf(line, "%ld", &ids[nids]);
        nids++;
    }

    int valid = 0;
    for(int i=0; i<nids; i++) {
        for(int j=0; j<nranges; j++) {
            if(ids[i] >= ranges[j].start && ids[i] <= ranges[j].end) {
                valid++;
                break;
            }
        }
    }

    printf("nranges: %d\n", nranges);
    printf("nids: %d\n", nids);
    printf("Valid: %d\n", valid);

    return 0;
}
