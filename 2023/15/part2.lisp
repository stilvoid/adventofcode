(ql:quickload "str")

(defvar *input*
  (with-open-file (file "input")
    (loop for line = (read-line file nil)
          while line
          collect line)))

(setf *input* (str:split "," (first *input*)))

(defvar *boxes* (make-hash-table))

(defun hash (s)
  (let ((cur 0))
    (loop for c across s do
          (setf cur (mod (* (+ cur (char-code c)) 17) 256)))
    cur))

(defun do-remove (s)
  (let ((boxnum (hash s)))
    (setf (gethash boxnum *boxes*) (remove s (gethash boxnum *boxes*) :key 'first :test 'string=))))

(defun do-put (s n)
  (let* ((boxnum (hash s))
    (box (gethash boxnum *boxes*))
    (existing (position s box :key 'first :test 'string=))
    (new (list s n)))
    (if existing
      (setf (elt (gethash boxnum *boxes*) existing) new)
      (push new (gethash boxnum *boxes*)))))

(defun do-one (s)
  (let ((parts (str:split "=" s)))
    (if (= (length parts) 1)
      (do-remove (subseq s 0 (- (length s) 1)))
      (do-put (first parts) (parse-integer (second parts))))))

(mapcar 'do-one *input*)

(defun count-box (boxnum)
  (let ((box (gethash boxnum *boxes*)))
    (reduce '+ (loop for item in box for i from 0 collect
        (* (+ boxnum 1) (- (length box) i) (second item))))))

(loop for i from 0 below 256 do
      (let ((box (gethash i *boxes*)))
        (when box (print (list i box)))))

(print (reduce '+ (mapcar 'count-box
              (remove nil
                      (loop for i from 0 below 256 collect
                            (when (gethash i *boxes*) i))))))
