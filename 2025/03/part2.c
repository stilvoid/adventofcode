#include <stdio.h>
#include <string.h>

unsigned long calc(int size, char* parts) {
    int N = 12;

    char total[N];

    int last=-1;

    for(int pos=0; pos<12; pos++) {
        char max = 0;
        int newLast = 0;

        for(int i=last+1; i<size-N+pos+1; i++) {
            if(parts[i]>max) {
                max = parts[i];
                newLast = i;
            }
        }

        last = newLast;

        total[pos] = max;
    }

    unsigned long out;
    sscanf(total, "%lu", &out);

    return out;
}

int main(int argc, char* argv[]) {
    char line[256];
    unsigned long total=0;

    while(scanf("%s\n", line) == 1) {
        total += calc(strlen(line), line);
    }

    printf("Total: %lu\n", total);

    return 0;
}
