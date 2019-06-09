package network

type MessageParser interface {
	Write(conn *TCPConn, args ...[]byte) error
	Read(conn *TCPConn) ([]byte, error)
	Conn(conn *TCPConn)
	Close(conn *TCPConn)
}
