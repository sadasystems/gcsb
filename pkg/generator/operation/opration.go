package operation

/*
 * Here we will be doing an 'operation generator'
 *
 * An operation generator returns a SQL operation (READ|WRITE)
 * with a configurable 'weighted probability'
 *
 * For example, I want to generate operations with 80% reads and 20% writes
 */

type Operation int

const (
	READ Operation = 1 + iota
	WRITE
)
