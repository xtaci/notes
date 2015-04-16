#include <stdio.h>
#include <stdlib.h>

struct Node {
	int data;
	struct Node * next;
};

struct Node * create_list();
void print_list(struct Node * head);
int lastnth(struct Node * head, size_t n);

int 
main(void) {
	struct Node * list = create_list();
	print_list(list);
	printf("3rd :%d\n", lastnth(list,3));
	printf("5th: %d\n", lastnth(list,5));
}

int 
lastnth(struct Node * head, size_t n) {
	struct Node * p = head, *p2 = head;

	for (int i=0;i<n-1;i++) {
		p = p->next;
	}

	while(p->next!=NULL)	{
		p=p->next;
		p2=p2->next;
	}

	return p2->data;
}
void 
print_list(struct Node * head) {
	printf("head->");
	for (struct Node *p = head;p!=NULL; p=p->next) {
		printf("%d->", p->data);
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
