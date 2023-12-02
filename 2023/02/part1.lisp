(defvar input (with-open-file (file "input")
  (loop for line = (read-line file nil)
        while line
        collect line)))

(defun split-one (delim text)
  (let ((pos (search delim text)))
    (if pos
      (cons (subseq text 0 pos)  (subseq text (+ pos (length delim))))
      (cons text NIL))))

(defun split (delim text)
  (let ((out (split-one delim text)))
    (if (cdr out)
      (append (list (car out)) (split delim (cdr out)))
      (list (car out)))))

(defun parse-colour (text)
  (let ((colour (split " " text)))
    (cons (elt colour 1) (parse-integer (car colour)))))

(defun parse-turn (text)
  (let ((colours (split ", " text)))
    (mapcar 'parse-colour colours)))

(defun parse-turns (text)
  (let ((turns (split "; " text)))
    (mapcar 'parse-turn turns)))

(defun parse-game (text)
  (let ((game (split ": " text)))
    (cons (parse-integer (elt (split " " (car game)) 1)) (parse-turns (elt game 1)))))

(defvar games (mapcar 'parse-game input))

(defvar quota (list
  (cons "red" 12)
  (cons "green" 13)
  (cons "blue" 14)))

(defun get-colour (colour colours)
  (assoc colour colours :test 'string=))

(defun possible-colour (colour)
  (<= (cdr colour) (cdr (get-colour (car colour) quota))))

(defun possible-turn (turn)
  (every 'possible-colour turn))

(defun possible-game (game)
  (every 'possible-turn (cdr game)))

(print (apply '+ (mapcar 'car (remove-if-not 'possible-game games))))
