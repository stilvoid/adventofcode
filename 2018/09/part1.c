#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#define PLAYERS 426
#define MARBLES 7205800

typedef struct node node;

struct node {
    unsigned int value;
    node* prev;
    node* next;
};

struct {
    node* pos;
    unsigned long int scores[PLAYERS];
} game;

void add(int value) {
    node* prev = game.pos;
    node* next = game.pos->next;

    node* newNode = calloc(sizeof(node), 1);
    newNode->value = value;
    newNode->prev = prev;
    newNode->next = next;

    prev->next = newNode;
    next->prev = newNode;

    game.pos = newNode;
}

void rem() {
    node* prev = game.pos->prev;
    node* next = game.pos->next;

    prev->next = next;
    next->prev = prev;

    free(game.pos);
    game.pos = next;
}

void play(int player, int value) {
    if(value % 23 == 0) {
        node* pos = game.pos;

        // Score points!
        for(int i=0; i<7; i++) {
            pos = pos->prev;
        }

        game.pos = pos;

        game.scores[player] += value + game.pos->value;

        rem();
    } else {
        game.pos = game.pos->next;
        add(value);
    }
}

int main() {
    int player = 0;
    unsigned long int winning_score = 0;

    for(int i=0; i<PLAYERS; i++) {
        game.scores[i] = 0;
    }

    game.pos = calloc(sizeof(node), 1);
    game.pos->value = 0;
    game.pos->prev = game.pos;
    game.pos->next = game.pos;

    for(int i=1; i<=MARBLES; i++) {
        play(player, i);

        player = (player + 1) % PLAYERS;
    }

    // Now find the best score
    for(int i=0; i<PLAYERS; i++) {
        if(game.scores[i] > winning_score) {
            winning_score = game.scores[i];
        }
    }
    printf("%lu\n", winning_score);
}
