(ql:quickload "str")

(defvar *input* (with-open-file (file "input")
                 (loop for line = (read-line file nil)
                       while line
                       collect line)))

(defvar *games* (mapcar (lambda (line) (let ((parts (str:split " " line)))
                                         (list :hand (first parts) :bid (parse-integer (second parts))))) *input*))

(print *games*)

(defconstant *card-rank* "AKQT98765432J")

(defun count-cards (hand)
  (loop for card across (remove-duplicates hand) collect
    (cons card (count card hand :test 'string=))))

(defun rank-card (a b)
  (if (= (cdr a) (cdr b)) (< (position (car a) *card-rank*) (position (car b) *card-rank*)) (> (cdr a) (cdr b))))

(defun best-card (hand)
  (if (= (length (remove #\J hand)) 0)
    (elt *card-rank* 0)
    (car (first (sort (count-cards (remove #\J hand)) 'rank-card)))))

(loop for game in *games* do
  (let* ((hand (getf game :hand)) (best (best-card hand)))
    (setf (getf game :pretend-hand) (if best
                                      (substitute (best-card hand) #\J hand)
                                      hand))))

(print *games*)

(defun score-counts (counts)
  (cond
    ((find 5 counts) 1)
    ((find 4 counts) 2)
    ((find 3 counts) (if (find 2 counts) 3 4))
    ((find 2 counts) (if (= (length counts) 3) 5 6))
    (T 7)))

(loop for game in *games* do
  (let* ((hand (getf game :pretend-hand))
         (card-counts (count-cards hand))
         (counts (loop for card-count in card-counts collect (cdr card-count)))
         (score (score-counts counts)))
    (setf (getf game :score) score)))


(defun hand< (a b)
  (loop for i from 0 to (- (length a) 1) do
    (let ((card-a (position (elt a i) *card-rank*)) (card-b (position (elt b i) *card-rank*)))
     (cond
       ((> card-a card-b) (return T))
       ((< card-a card-b) (return NIL))))))

(defun game< (a b)
  (let ((a-score (getf a :score)) (b-score (getf b :score)))
    (if (= a-score b-score)
      (hand< (getf a :hand) (getf b :hand))
      (> a-score b-score))))

(setf *games* (sort *games* 'game<))

(print *games*)

(print (apply '+ (loop for i from 1 to (length *games*) for game in *games* collect
                  (* i (getf game :bid)))))
