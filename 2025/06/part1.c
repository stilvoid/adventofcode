#include <stdio.h>
#include <string.h>

long long nums[10][1000];
char ops[1000];

char lines[100][5000];

int main() {
    int rows = 0;
    int cols = 0;

    while(fgets(lines[rows], sizeof(lines[rows]), stdin)) {
        rows++;
    }
    rows--;

    for(int row=0; row<rows; row++) {
        int col = 0;
        char* tok = strtok(lines[row], " ");
        while(tok != NULL) {
            if(sscanf(tok, "%lld", &nums[row][col]) ==1) {
                col++;
            }

            tok = strtok(NULL, " ");
        }
        
        if(cols == 0) {
            cols = col;
        }
    }

    char* tok = strtok(lines[rows], " ");
    int col = 0;
    while(tok != NULL) {
        if(sscanf(tok, "%c", &ops[col]) == 1) {
            col++;
        }

        tok = strtok(NULL, " ");
    }

    long long total = 0;
    for(int col=0; col<cols; col++) {
        long long n = nums[0][col];
        for(int row=1; row<rows; row++) {
            if(ops[col] == '+') {
                n += nums[row][col];
            } else {
                n *= nums[row][col];
            }
        }
        total += n;
    }

    printf("Total: %lld\n", total);
}
