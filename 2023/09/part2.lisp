(ql:quickload "str")

(defvar *input*
  (with-open-file (file "input")
    (loop for line = (read-line file nil)
          while line
          collect line)))

(setf *input* (loop for line in *input* collect
                    (mapcar 'parse-integer (str:split " " line))))

(defun diffs (seq)
  (loop for a in (subseq seq 0 (- (length seq) 1)) 
        for b in (subseq seq 1) collect
        (- b a)))

(defun get-next (seq)
  (if (apply '= seq)
    (first seq)
    (- (first seq) (get-next (diffs seq)))))

(print (apply '+ (mapcar 'get-next *input*)))
