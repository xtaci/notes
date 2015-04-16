#include <stdio.h>
#include <stdlib.h>

struct Node {
	int data;
	struct Node * next;
};

void remove_dup(struct Node * head);
struct Node * create_list();
void print_list(struct Node * head);

int 
main(void) {
	struct Node * list = create_list();
	print_list(list);
	remove_dup(list);
	print_list(list);
}

void 
print_list(struct Node * head) {
	printf("head->");
	for (struct Node *p = head;p!=NULL; p=p->next) {
		printf("%d->", p->data);
	}
	printf("NULL\n");
}

void
remove_dup(struct Node * head) {
	struct Node * p1;
	for (p1=head;p1!=NULL;p1=p1->next) {
		struct Node * p2, *prev = p1;
		for (p2=p1->next;p2!=NULL;p2=p2->next) {
			if (p2->data == p1->data) {
				prev->next = p2->next;
				free(p2);
			} else {
				prev = p2;
			}
		}
	}
}

struct Node * 
create_list() {
	struct Node * head = malloc(sizeof(struct Node));
	head->data = 0;
	head->next = NULL;
	struct Node *p = head;
	for(int i=1;i<100;i++) {
		p->next = malloc(sizeof(struct Node));
		p=p->next;
		p->data=i%10;
		p->next=NULL;
	}
	return head;
}
