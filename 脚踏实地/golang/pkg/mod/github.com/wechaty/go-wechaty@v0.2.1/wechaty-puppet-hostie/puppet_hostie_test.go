package puppethostie

// TODO 建议 mock http
//func TestPuppetHostie_discoverHostieIP(t *testing.T) {
//	type fields struct {
//		Puppet      *wechatyPuppet.Puppet
//		grpcConn    *grpc.ClientConn
//		grpcClient  wechaty.PuppetClient
//		eventStream wechaty.Puppet_EventClient
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		wantS   string
//		wantErr bool
//	}{
//		{
//			name: "0.0.0.0",
//			fields: fields{
//				Puppet: &wechatyPuppet.Puppet{
//					Option: &wechatyPuppet.Option{Token: "__TOKEN__"},
//				},
//			},
//			wantS:   "0.0.0.0",
//			wantErr: false,
//		},
//		{
//			name: "timeout",
//			fields: fields{
//				Puppet: &wechatyPuppet.Puppet{
//					Option: &wechatyPuppet.Option{Token: "__TOKEN__", Timeout: 1 * time.Nanosecond},
//				},
//			},
//			wantS:   "",
//			wantErr: true,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := &PuppetHostie{
//				Puppet:      tt.fields.Puppet,
//				grpcConn:    tt.fields.grpcConn,
//				grpcClient:  tt.fields.grpcClient,
//				eventStream: tt.fields.eventStream,
//			}
//			gotS, err := p.discoverHostieIP()
//			if (err != nil) != tt.wantErr {
//				t.Errorf("discoverHostieIP() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if gotS != tt.wantS {
//				t.Errorf("discoverHostieIP() gotS = %v, want %v", gotS, tt.wantS)
//			}
//		})
//	}
//}
