#include <stdio.h>
#include <stdlib.h>

typedef struct {
    long long x;
    long long y;
} Point;

typedef struct {
    Point* A;
    Point* B;
    long long area;
} Rect;

#define MAXPOINTS 500

Point points[MAXPOINTS];
Rect rects[MAXPOINTS * MAXPOINTS];

int cmpRects(const void *A, const void *B) {
    const Rect *a = A;
    const Rect *b = B;

    if(a->area < b->area) {
        return 1;
    } else if(a->area > b->area) {
        return -1;
    }
    return 0;
}

long long area(Point *A, Point *B) {
    long long dx = A->x - B->x;
    long long dy = A->y - B->y;

    if(dx<0) {
        dx=-dx;
    }

    if(dy<0) {
        dy=-dy;
    }

    return (dx+1)*(dy+1);
}

int main() {
    int npoints = 0;
    int nrects = 0;

    while(scanf("%lld,%lld", &points[npoints].x, &points[npoints].y) == 2) {
        npoints++;
    }

    for(int i=0; i<npoints-1; i++) {
        for(int j=i+1; j<npoints; j++) {
            Point *A = &points[i];
            Point *B = &points[j];

            rects[nrects].A = A;
            rects[nrects].B = B;

            rects[nrects].area = area(A, B);

            nrects++;
        }
    }

    qsort(rects, nrects, sizeof(rects[0]), cmpRects);

    printf("Biggest: %lld\n", rects[0].area);
}
