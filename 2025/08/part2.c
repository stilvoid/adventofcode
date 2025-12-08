#include <stdio.h>
#include <stdlib.h>

typedef struct {
    long long x;
    long long y;
    long long z;
    int circuit;
} Box;

typedef struct {
    Box* a;
    Box* b;
    long long dist;
} Pair;

Box boxes[1000];
Pair pairs[1000 * 1000];
int circuits[1000];

int cmpPair(const void *a, const void *b) {
    const Pair *A = a;
    const Pair *B = b;

    if(A->dist < B->dist) {
        return -1;
    } else if(A->dist > B->dist) {
        return 1;
    }

    return 0;
}

int cmpInt(const void *a, const void *b) {
    const int *A = a;
    const int *B = b;

    return *B - *A;
}

int main() {
    int ncircuits = 0;
    int nboxes = 0;
    int npairs = 0;

    while(scanf("%lld,%lld,%lld", &boxes[nboxes].x, &boxes[nboxes].y, &boxes[nboxes].z) == 3) {
        boxes[nboxes].circuit = -1;
        nboxes++;
    }

    for(int i=0; i<nboxes-1; i++) {
        for(int j=i+1; j<nboxes; j++) {
            pairs[npairs].a = &boxes[i];
            pairs[npairs].b = &boxes[j];

            long long dx = boxes[i].x - boxes[j].x;
            long long dy = boxes[i].y - boxes[j].y;
            long long dz = boxes[i].z - boxes[j].z;

            pairs[npairs].dist = (dx*dx) + (dy*dy) + (dz*dz);

            npairs++;
        }
    }

    qsort(pairs, npairs, sizeof(pairs[0]), cmpPair);

    // Number of circuits
    int count = nboxes;

    int i;

    // Connect pairs
    for(i=0; i<npairs; i++) {
        int aCirc = pairs[i].a->circuit;
        int bCirc = pairs[i].b->circuit;

        if(aCirc >= 0) {
            if(bCirc >= 0) {
                if(aCirc != bCirc) {
                    for(int j=0; j<nboxes; j++) {
                        if(boxes[j].circuit == bCirc) {
                            boxes[j].circuit = aCirc;
                        }
                    }
                    count--;
                }
            } else {
                pairs[i].b->circuit = aCirc;
                count--;
            }
        } else if(pairs[i].b->circuit >= 0) {
            pairs[i].a->circuit = bCirc;
            count--;
        } else {
            pairs[i].a->circuit = ncircuits;
            pairs[i].b->circuit = ncircuits;
            ncircuits++;
            count--;
        }

        if(count == 1) {
            break;
        }
    }

    printf("Answer: %lld\n", pairs[i].a->x * pairs[i].b->x);

    return 0;
}
