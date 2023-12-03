(defvar *input* (with-open-file (file "input")
  (loop for line = (read-line file nil)
        while line
        collect line)))

(setf *input* (mapcar (lambda (line) (loop for c across line collect c)) *input*))

(defvar *rows* (length *input*))

(defvar *cols* (length (car *input*)))

(defun getl (row col)
  (if (and (>= row 0) (< row *rows*) (>= col 0) (< col *cols*))
    (elt (elt *input* row) col)
    NIL))

(defun setl (row col val)
  (if (and (>= row 0) (< row *rows*) (>= col 0) (< col *cols*))
      (setf (elt (elt *input* row) col) val)))

(loop for line in *input* do (print (concatenate 'string line)))

(defun read-num (row startcol)
  (let ((out (list (getl row startcol))))
    (loop for col from (+ startcol 1) to *cols* do
      (let ((c (getl row col)))
        (print (list col c))
        (if (and c (digit-char-p c)) (nconc out (list c)) (return out))))))

; Find a number, record position of gears next to it
(defun find-num-in-row (row)
  (remove nil (loop for col from 0 to (- *cols* 1) collect
    (let ((c (getl row col)))
      (if (digit-char-p c) (let ((num (read-num row col)) (ocol col))
        (setf col (+ col (length num) -1))
        (find-gears num row ocol)))))))

(defun find-gears (num startrow startcol)
  (print (list "OO:" num startrow startcol))
  (cons num (remove nil (loop for row from (- startrow 1) to (+ startrow 1) collect
    (remove nil (loop for col from (- startcol 1) to (+ startcol (length num)) collect
      (if (equal (getl row col) #\*) (cons row col) NIL)))))))

(defvar *nums* (mapcan 'find-num-in-row (loop for row from 0 to (- *rows* 1) collect row)))

(print "NUMS:")
(loop for num in *nums* do (print num))

(defvar *gears* (make-hash-table :test 'equal))

(defun contains-gear (num gear)
  (if (= (length (cdr num)) 0) NIL
  (some (lambda (g) (equal gear g)) (cdr num))))

(defun check-gear (gear)
  (setf (gethash gear *gears*) (mapcar 'car (remove-if-not (lambda (num) (contains-gear num gear)) *nums*))))

(defun check-gears (num)
  (mapcar 'check-gear (cdr num)))

(mapcar 'check-gears *nums*)

(print "GEARS:")
(maphash (lambda (gear nums) (format t "~a: ~a~%" gear nums)) *gears*)

(maphash (lambda (gear nums) (setf (gethash gear *gears*) (mapcar (lambda (num) (parse-integer (concatenate 'string num))) nums))) *gears*)

(print *gears*)

(print (apply '+ (remove nil (loop for v being the hash-values of *gears* collect (if (= (length v) 2) (apply '* v) NIL)))))
