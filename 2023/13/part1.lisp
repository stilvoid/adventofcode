(ql:quickload "str")

(defvar *input*
  (with-open-file (file "input")
    (loop for line = (read-line file nil)
          while line
          collect line)))

(defun split (seq)
  (let ((out nil) (cur nil))
    (loop for line in seq do
          (if (string= line "")
            (progn
              (setf cur (nreverse cur))
              (push cur out)
              (setf cur nil))
            (push line cur)))
    (setf cur (nreverse cur))
    (push cur out)
    (nreverse out)))

(defun rotate (seq)
  (let ((out nil))
    (loop for i from 0 below (length (first seq)) collect
          (concatenate 'string (reverse (loop for j from 0 below (length seq) collect
                (char (elt seq j) i)))))))

(defun reflect-p (a b)
  "a and b are split at a possible point of reflection"
  (every 'equal b (reverse a)))

(defun find-reflect (seq)
  (loop for i from 1 below (length seq) do
       (when (reflect-p (subseq seq 0 i) (subseq seq i)) (return i))))

(defun do-record (seq)
  (let ((horiz (find-reflect seq)))
    (if horiz
      (* horiz 100)
      (find-reflect (rotate seq)))))

(setf *input* (split *input*))

(print (reduce '+ (mapcar (lambda (record) (print (do-record record))) *input*)))
