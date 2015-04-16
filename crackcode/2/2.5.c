#include <stdio.h>
#include <stdlib.h>

struct Node {
	int data;
	struct Node * next;
};

struct Node * create_list();
void print_list(struct Node * head);
void loop_detect(struct Node * head);

int 
main(void) {
	loop_detect(create_list());
}

void 
loop_detect(struct Node * head) {
	struct Node * p1, *p2;
	p1 = p2 = head;

	while (p2->next != NULL && p2->next->next != NULL) {
		p1 = p1->next;
		p2 = p2->next->next;
		if (p1 == p2) {
			break;
		}
	}

	if (p2->next == NULL || p2->next->next == NULL) {
		printf("no loop");
		return;
	}

	p1 = head;
	while(p1 !=p2) {
		p1 = p1->next;
		p2 = p2->next;
	}

	printf("loop point (%d, %p)\n", p1->data, p1);
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
	struct Node *loop_point;
	for(int i=1;i<10;i++) {
		p->next = malloc(sizeof(struct Node));
		p=p->next;
		p->data=i%10;
		p->next=NULL;
		if (i==5) {
			loop_point = p;
		}
	}

	p->next = loop_point;

	return head;
}
