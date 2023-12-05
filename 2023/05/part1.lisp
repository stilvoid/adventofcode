(ql:quickload "str")

(defvar *input* (with-open-file (file "input")
  (loop for line = (read-line file nil)
        while line
        collect line)))

(defvar *seeds* (mapcar 'parse-integer (str:split " " (second (str:split ": " (first *input*))))))

(defun parse-map (seq)
  (loop for line in (subseq seq 1) until (string= line "") collect
    (mapcar 'parse-integer (str:split " " line))))

(defun parse-maps (seq)
  (let ((out ()))
    (loop until (= (length seq) 0) collect
      (let ((next (parse-map seq)))
        (push next out)
        (if (> (length seq) (+ (length next) 1))
          (setf seq (subseq seq (+ 2 (length next))))
          (setf seq ()))))
    (nreverse out)))

(defvar *mappers* (parse-maps (subseq *input* 2)))

(defun map-to-next (mapper value)
  (loop for entry in mapper do
    (when (and (>= value (second entry)) (< value (+ (second entry) (third entry))))
      (return-from map-to-next (+ (first entry) (- value (second entry))))))
  value)

(defun map-through (value)
  (let ((cur value))
    (loop for mapper in *mappers* do
      (setf cur (map-to-next mapper cur)))
    cur))

(print (apply 'min (loop for seed in *seeds* collect (map-through seed))))
