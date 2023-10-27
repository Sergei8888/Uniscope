package Earth

//func TestNewLocation(t *testing.T) {
//	cases := []struct {
//		input    string
//		expected *Earth.Earth
//	}{
//		{"118 20 17 W 33 50 41 N", &Earth.Earth{
//			latitudeDeg:  33,
//			latitudeMin:  50,
//			latitudeSec:  41,
//			longitudeDeg: 118,
//			longitudeMin: 20,
//			longitudeSec: 17,
//			fromSouth:    false,
//			fromWest:     true,
//		}},
//	}
//
//	for i, tc := range cases {
//		t.Run(fmt.Sprintf("test #%d", i+1), func(t *testing.T) {
//			loc := NewLocation(tc.input)
//			assert.Equal(t, tc.expected, loc)
//		})
//	}
//}
