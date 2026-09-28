// mautrix-whatsapp - A Matrix-WhatsApp puppeting bridge.
// Copyright (C) 2026 Tulir Asokan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package msgconv

import (
	"errors"
	"testing"

	"go.mau.fi/whatsmeow/proto/waE2E"
	"maunium.net/go/mautrix/bridgev2"
)

func TestWrapViewOnce(t *testing.T) {
	img := &waE2E.ImageMessage{}
	wrapped, err := wrapViewOnce(&waE2E.Message{ImageMessage: img})
	if err != nil || wrapped.GetViewOnceMessageV2().GetMessage().GetImageMessage() != img || !img.GetViewOnce() {
		t.Errorf("image wasn't wrapped as view-once: %v %v", wrapped, err)
	}
	vid := &waE2E.VideoMessage{}
	wrapped, err = wrapViewOnce(&waE2E.Message{VideoMessage: vid})
	if err != nil || wrapped.GetViewOnceMessageV2().GetMessage().GetVideoMessage() != vid || !vid.GetViewOnce() {
		t.Errorf("video wasn't wrapped as view-once: %v %v", wrapped, err)
	}
	_, err = wrapViewOnce(&waE2E.Message{DocumentMessage: &waE2E.DocumentMessage{}})
	if !errors.Is(err, bridgev2.ErrViewOnceNotAllowed) {
		t.Errorf("expected ErrViewOnceNotAllowed for document, got %v", err)
	}
}
