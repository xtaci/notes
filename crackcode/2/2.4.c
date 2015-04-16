#include <stdio.h>
#include <stdlib.h>

struct Node {
	int data;
	struct Node * next;
};

struct Node * create_list();
void print_list(struct Node * head);
void list_sum(struct Node * a, struct Node * b);

int 
main(void) {
	list_sum(create_list(), create_list());
}

void 
list_sum(struct Node * a, struct Node * b) {
	int carry = 0;
	for (;;) {
		int sum = a->data + b->data + carry;
		printf("%d->", sum%10);
		carry = sum/10;

		a = a->next;b=b->next;
		if (a== NULL || b == NULL) {
			break;
		}
	}

	printf("%d\n", carry);
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
