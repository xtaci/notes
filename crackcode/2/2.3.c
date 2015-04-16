#include <stdio.h>
#include <stdlib.h>

struct Node {
	int data;
	struct Node * next;
};

struct Node * create_list();
void print_list(struct Node * head);
void delete_middle(struct Node * n);

int 
main(void) {
	struct Node * list = create_list();
	struct Node * p = list;
	for (int i=0;i<2;i++) {
		p = p->next;
	}

	print_list(list);
	printf("delete %d %p\n", p->data, p);
	delete_middle(p);
	print_list(list);
}

void
delete_middle(struct Node * n) {
	struct Node * next= n->next;
	n->next = next->next;
	n->data = next->data;
	free(next);
}

void 
print_list(struct Node * head) {
	printf("head->");
	for (struct Node *p = head;p!=NULL; p=p->next) {
		printf("(%d %p)->", p->data, p);
	}
	printf("NULL\n");
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
