#include <stdio.h>
#include <stdlib.h>

typedef struct node node;

struct node {
  int data;
  node *previous;
  node *next;
};

typedef struct {
  node *head;
  node *tail;
  int size;
} linked;

linked *createLinked() {
  linked *l = malloc(sizeof(linked));
  l->head = NULL;
  l->tail = NULL;
  l->size = 0;

  return l;
}

void pushFront(linked *l, int data) {
  node *n = malloc(sizeof(node));
  n->data = data;
  n->previous = NULL;
  n->next = l->head;

  if (l->head != NULL) {
    l->head->previous = n;
  } else {
    l->tail = n;
  }

  l->head = n;
  l->size++;
}

void pushBack(linked *l, int data) {
  node *n = malloc(sizeof(node));
  n->data = data;
  n->previous = l->tail;
  n->next = NULL;

  if (l->tail != NULL) {
    l->tail->next = n;
  } else {
    l->head = n;
  }

  l->tail = n;
  l->size++;
}

void popBack(linked *l) {
  if (l->tail == NULL)
    return;

  if (l->tail->previous == NULL) {
    l->head = NULL;
    l->tail = NULL;
  } else {
    l->tail = l->tail->previous;
    l->tail->next = NULL;
  }

  l->size--;
}

void popFront(linked *l) {
  if (l->head == NULL)
    return;

  if (l->head->next == NULL) {
    l->head = NULL;
    l->tail = NULL;
  } else {
    l->head = l->head->next;
    l->head->previous = NULL;
  }

  l->size--;
}

void freeLinked(linked *l) {
  node *current = l->head;
  while (current != NULL) {
    node *next = current->next;
    free(current);
    current = next;
  }

  free(l);
}

node *search(linked *l, int data) {
  node *current = l->head;
  while (current->data != data) {
    if (current->next == NULL) {
      return NULL;
    }
    current = current->next;
  }

  return current;
}

void printl(linked *l) {
  node *current = l->head;
  while (current != NULL) {
    printf("%d ", current->data);
    current = current->next;
  }
  printf("\n");
}

void printn(node *n) {
  if (n->previous)
    printf("prev:%d, ", n->previous->data);
  else
    printf("prev:NULL, ");

  printf("data:%d, ", n->data);

  if (n->next)
    printf("next:%d\n", n->next->data);
  else
    printf("next:NULL\n");
}

int main() {
  linked *linked = createLinked();
  pushBack(linked, 18);
  pushBack(linked, 3);
  pushBack(linked, 2);
  pushBack(linked, 10);

  popBack(linked);
  popFront(linked);

  printl(linked);

  node *searchBy2 = search(linked, 2);
  printn(searchBy2);
}
