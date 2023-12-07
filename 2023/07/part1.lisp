(ql:quickload "str")

(defvar *input* (with-open-file (file "input")
                 (loop for line = (read-line file nil)
                       while line
                       collect line)))

(defvar *games* (mapcar (lambda (line) (let ((parts (str:split " " line)))
                                         (list :hand (first parts) :bid (parse-integer (second parts))))) *input*))

(defconstant *card-rank* "AKQJT98765432")

(defun count-cards (original-hand)
  (let ((hand (copy-seq original-hand)))
    (loop while (> (length hand) 0) collect
      (let ((card (elt hand 0)))
       (setf num (count card hand :test 'string=))
       (setf hand (remove card hand))
       num))))

(defun score-counts (counts)
  (cond
    ((find 5 counts) 1)
    ((find 4 counts) 2)
    ((find 3 counts) (if (find 2 counts) 3 4))
    ((find 2 counts) (if (= (length counts) 3) 5 6))
    (T 7)))

(print *games*)

(loop for game in *games* do
  (let* ((hand (getf game :hand)) (counts (count-cards hand)) (score (score-counts counts)))
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
