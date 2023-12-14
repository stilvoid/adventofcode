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

(defun find-reflects (seq)
  (remove nil (loop for i from 1 below (length seq) collect
       (when (reflect-p (subseq seq 0 i) (subseq seq i)) i))))


(defun do-record (seq)
  (let ((horiz (find-reflects seq)) (vert (find-reflects (rotate seq))) (out NIL))
    (mapcar (lambda (n) (push (* n 100) out)) horiz)
    (mapcar (lambda (n) (push n out)) vert)
    out))

(defun dup (seq) (mapcar 'copy-seq seq))

(defun find-smudge (seq)
  (let ((old (first (do-record seq))))
    (first (remove old (reduce 'append (loop for y from 0 below (length seq) collect
          (reduce 'append (loop for x from 0 below (length (first seq)) collect
                (let ((new-seq (dup seq)) (c (char (elt seq y) x)))
                  (setf (char (elt new-seq y) x) (if (char= c #\.) #\# #\.))
                  (do-record new-seq))))))))))

(setf *input* (split *input*))

(print (reduce '+ (mapcar (lambda (record) (print (find-smudge record))) *input*)))
