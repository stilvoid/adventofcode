#include <stdio.h>
#include <string.h>

char nums[10][1000][20];
char ops[1000];

char lines[100][5000];

int main() {
    int rows = 0;
    int len = 0;
    while(fgets(lines[rows], sizeof(lines[rows]), stdin)) {
        int l = strlen(lines[rows]);
        if(l > len) {
            len = l;
        }
        rows++;
    }
    rows--;

    int start = 0;
    int end = 0;

    long long total = 0;

    for(int i=1; i<len; i++) {
        if(lines[rows][i] != ' ') {
            if(lines[rows][i] == 0) {
                end = len - 1;
                i = len;
            } else {
                end = i;
            }

            long long col_total = 0;

            // Do stuff
            for(int pos=start; pos<end; pos++) {
                long long n = 0;
                for(int row=0; row<rows; row++) {
                    if(lines[row][pos] != ' ') {
                        n *= 10;
                        n += lines[row][pos] - '0';
                    }
                }

                if(n != 0) {
                    if(col_total == 0) {
                        col_total = n;
                    } else {
                        if(lines[rows][start] == '+') {
                            col_total += n;
                        } else {
                            col_total *= n;
                        }
                    }
                }
            }

            total += col_total;
            
            // Get ready for next
            start = i;
        }
    }

    printf("Total: %lld\n", total);
}
