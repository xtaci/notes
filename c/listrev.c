#include <stdio.h>
#include <stdlib.h>

struct Node {
	int data;
	struct Node * next;
};

struct Node * create_list();
void print_list(struct Node * head);
struct Node * listrev(struct Node * head);

int
main(void) {
	struct Node * list = create_list();
	print_list(list);
	print_list(listrev(list));
}

struct Node *
listrev(struct Node * p) {
	if (p==NULL) {
		return p;
	}

	struct Node * tmp;
	struct Node * head = NULL;

	while (p != NULL) {
		tmp = p->next;
		p->next = head;
		head = p;
		p = tmp;
	}
	return head;
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
