(ql:quickload "str")

(defvar *input*
  (with-open-file (file "input")
    (loop for line = (read-line file nil)
          while line
          collect line)))

(defconstant *steps* (first *input*))

(defun parse-directions (directions)
  (let ((parts (str:split ", " directions)))
    (list (subseq (first parts) 1) (subseq (second parts) 0 3))))

(defun parse-instruction (instruction)
  (let ((parts (str:split " = " instruction)))
    (list (first parts) (parse-directions (second parts)))))

(defparameter *network* (mapcar 'parse-instruction (subseq *input* 2)))

(print *network*)

(defconstant L #\L)
(defconstant R #\R)

(defun get-next (from direction)
  (let ((opts (second (assoc from *network* :test 'string=))))
    (if (char= direction L) (first opts) (second opts))))

(defun run (start)
  (let ((pos start))
    (loop for i from 0 do
          (let ((dir (elt *steps* (mod i (length *steps*)))))
            (print (list pos dir))
            (setf pos (get-next pos dir))
            (when (string= pos "ZZZ") (return (+ i 1)))))))

(print (run "AAA"))
