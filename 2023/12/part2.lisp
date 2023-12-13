(ql:quickload "str")

(defvar *input*
  (with-open-file (file "input")
    (loop for line = (read-line file nil)
          while line
          collect line)))

(defun split-chunk (chunk)
  (remove "" (str:split "." chunk) :test 'string=))

(defun count-chunk (chunk)
  (mapcar 'length (split-chunk chunk)))

(setf *input* (mapcar (lambda (line)
                        (let* ((parts (str:split " " line))
                               (chunks (split-chunk (first parts)))
                               (counts (mapcar 'parse-integer (str:split "," (second parts)))))
                          (list :chunks chunks :counts counts))) *input*))

(defun get-perms (chunk)
  (if (= (length chunk) 1)
    (if (string= chunk "?")
      (list "#" ".")
      (list chunk))
    (let ((c (char chunk 0))
          (next (get-perms (subseq chunk 1))))
      (if (char= c #\?)
        (append
          (loop for opt in next collect
                (concatenate 'string "#" opt))
          (loop for opt in next collect
                (concatenate 'string "." opt)))
        (loop for opt in next collect
              (concatenate 'string (subseq chunk 0 1) opt))))))

(defun do-chunk (chunk)
  (loop for perm in (get-perms chunk) collect
        (list :chunk perm :counts (count-chunk perm))))

(defun do-record (record)
  (let ((n 0)
        (chunks (getf record :chunks))
        (counts (getf record :counts)))
    (when (and (null chunks) (null counts)) (return-from do-record (list "")))
    (when (and (null chunks) (not (null counts))) (return-from do-record NIL))
    (reduce 'append (loop for result in (do-chunk (first chunks)) collect
          (when (and (<= (length (getf result :counts)) (length counts)) (every '= (getf result :counts) counts))
            (loop for next in (do-record (list :chunks (subseq chunks 1) :counts (subseq counts (length (getf result :counts))))) collect
                  (concatenate 'string (getf result :chunk) next)))))))

(print (reduce '+ (mapcar (lambda (record) (length (do-record record))) *input*)))
