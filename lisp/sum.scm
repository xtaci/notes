(define (sum term a next b)
	(if (> a b)
		0
		(+ (term a) (sum term (next a) next b))
	)
)

(define (sum-int a b)
	(define (term n) n)
	(define (next n) (+ n 1))
	(sum term a next b)
)

(define (sum-square a b)
	(define (term n) (* n n))
	(define (next n) (+ n 1))
	(sum term a next b)
)
