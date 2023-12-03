(defvar input (with-open-file (file "input")
  (loop for line = (read-line file nil)
        while line
        collect line)))

(setf input (mapcar (lambda (line) (loop for c across line collect c)) input))

(loop for line in input do (print line))

(defvar blessings (mapcar (lambda (line) (loop for c in line collect NIL)) input))

(defun getl (l row col)
  (if (and (>= row 0) (< row (length l)) (>= col 0) (< col (length (car l))))
    (elt (elt l row) col)
    NIL))

(defun setl (l row col val)
  (if (and (>= row 0) (< row (length l)) (>= col 0) (< col (length (car l))))
      (setf (elt (elt l row) col) val)))

(defun bless (row col)
  (print (list "blessing" row col))
  (setl blessings row col T)
  (let ((r1 (max 0 (- row 1)))
    (r2 (min (length input) (+ row 1)))
    (c1 (max 0 (- col 1)))
    (c2 (min (length (car input)) (+ col 1))))

    (loop for r from r1 to r2 do
      (loop for c from c1 to c2 do
        (if (and (not (equal #\. (getl input r c)))
          (getl input r c)
          (not (getl blessings r c))) (bless r c))))))

(loop for row from 0 to (- (length input) 1) do
  (loop for col from 0 to (- (length (car input)) 1) do
    (let ((c (getl input row col)))
      (if (and (not (equal c #\.)) (not (digit-char-p c))) (bless row col)))))

(loop for line in blessings do (print line))

(loop for row from 0 to (- (length input) 1) do
  (loop for col from 0 to (- (length (car input)) 1) do
    (let ((b (getl blessings row col)) (c (getl input row col)))
      (if (or (not b) (not (digit-char-p c))) (setl input row col #\ )))))

(loop for line in input do (print line))

(print (apply '+ (remove nil (mapcan (lambda (line) (with-input-from-string (in (concatenate 'string line))
  (loop for n = (read in nil nil) while n collect n))) input))))
