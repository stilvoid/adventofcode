(defvar input (with-open-file (file "input")
  (loop for line = (read-line file nil)
        while line
        collect line)))

(defvar terms '(
    ("one" . 1)
    ("two" . 2)
    ("three" . 3)
    ("four" . 4)
    ("five" . 5)
    ("six" . 6)
    ("seven" . 7)
    ("eight" . 8)
    ("nine" . 9)
    ("0" . 0)
    ("1" . 1)
    ("2" . 2)
    ("3" . 3)
    ("4" . 4)
    ("5" . 5)
    ("6" . 6)
    ("7" . 7)
    ("8" . 8)
    ("9" . 9)))

(defun find-from-left (text)
  (loop for term in terms
    do (if (equal 0 (search (car term) text)) (return (cdr term)))))

(defun find-all-from-left (text)
  (remove nil
    (loop for i from 0 to (- (length text) 1)
      collect (find-from-left (subseq text i)))))

(defun first-and-last (nums)
  (+ (* 10 (car nums)) (car (last nums))))

(print (apply '+ (loop for row in input collect (first-and-last (find-all-from-left row)))))
