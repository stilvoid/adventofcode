#include <stdio.h>

typedef struct {
    long start;
    long end;
} Range;

Range ranges[2000];

int main(int argc, char* argv[]) {
    char line[100];

    int nranges = 0;

    while(fgets(line, 100, stdin)) {
        if(line[0] == '\n') {
            break;
        }

        sscanf(line, "%ld-%ld", &ranges[nranges].start, &ranges[nranges].end);
        nranges++;
    }

    // Fix dupes
    for(int i=0; i<nranges-1; i++) {
        Range* a = &ranges[i];

        for(int j=i+1; j<nranges; j++) {
            Range* b = &ranges[j];

            if(a->end < b->start || a->start > b->end) {
                continue;
            }

            if(a->start >= b->start && a->end <= b->end) {
                // Fully inside
                a->start = -1;
                a->end = -1;
            } else if(a->start < b->start && a->end <= b->end) {
                // Left overlap
                a->end = b->start - 1;
            } else if(a->start >= b->start && a->end > b->end) {
                // Right overlap
                a->start = b->end + 1;
            } else if(a->start < b->start && a->end > b->end) {
                // Overhang both ends
                b->start = -1;
                b->end = -1;
            }
        }
    }

    long total = 0;
    for(int i=0; i<nranges; i++) {
        if(ranges[i].start != -1) {
            total += ranges[i].end - ranges[i].start + 1;
        }
    }

    printf("Total: %ld\n", total);

    return 0;
}
