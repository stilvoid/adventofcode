#include <stdio.h>
#include <string.h>

int calc(int size, char* parts) {
    int max=0;

    for(int i=0; i<size-1; i++) {
        for(int j=i+1; j<size; j++) {
            int n = (parts[i] - '0') * 10 + parts[j] - '0';
            if(n>max) {
                max = n;
            }
        }
    }

    return max;
}

int main(int argc, char* argv[]) {
    char line[256];
    int total=0;

    while(scanf("%s\n", line) == 1) {
        total += calc(strlen(line), line);
    }

    printf("Total: %d\n", total);

    return 0;
}
