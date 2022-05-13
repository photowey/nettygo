/*
 * Copyright © 2022 photowey (photowey@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package session

import (
	"net"

	"github.com/photowey/nettygo/interal/remoting/connection"
)

type Session interface {
	connection.Connection
	Reset()
	Conn() net.Conn
	Stat() string
	IsClosed() bool
}

type Reader interface {
	Read(Session, []byte) (any, int, error) // buf.ByteBuf?
}

type Writer interface {
	Write(Session, any) ([]byte, error)
}

type ReadWriter interface {
	Reader
	Writer
}

type EventListener interface {
	OnOpen(Session) error
	OnClose(Session)
	OnMessage(Session, any)
	OnError(Session, error)
}
