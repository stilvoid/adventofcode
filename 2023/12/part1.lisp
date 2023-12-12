(ql:quickload "str")

(defvar *input*
  (with-open-file (file "input")
    (loop for line = (read-line file nil)
          while line
          collect line)))

(setf *input* (mapcar (lambda (line)
                        (let* ((parts (str:split " " line))
                               (data (first parts))
                               (check (mapcar 'parse-integer (str:split "," (second parts)))))
                          (list :data data :check check))) *input*))

(defun check (data)
  (mapcar 'length (remove "" (str:split "." data) :test 'string=)))

(defun verify (record)
  (equal (check (getf record :data)) (getf record :check)))

(defun get-perms (n)
  (if (= n 1)
    (list (list #\.) (list #\#))
    (append
      (loop for result in (get-perms (- n 1)) collect
          (append (list #\.) result))
      (loop for result in (get-perms (- n 1)) collect
          (append (list #\#) result)))))

(defun find-qs (data)
  (remove nil (loop for c across data for pos from 0 collect
                    (when (char= c #\?) pos))))

(defun munge-perm (data locs perm)
  (let ((out (copy-seq data)))
    (loop for loc in locs for c in perm do
          (setf (char out loc) c))
    out))

(defun get-options (record)
  (let* ((data (getf record :data))
        (locs (find-qs data))
        (perms (get-perms (length locs))))
    (loop for perm in perms collect
          (munge-perm data locs perm))))

(defun get-valid-options (record)
  (remove nil (loop for opt in (get-options record) collect
        (when (verify (list :data opt :check (getf record :check))) opt))))

(print (reduce '+ (mapcar (lambda (record) (length (get-valid-options record))) *input*)))
