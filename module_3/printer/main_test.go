package printer

import "testing"


func TestGetUTFLength(t *testing.T) {
        cases := []struct {
            name string
            values []byte
            want []interface{}
        }{
            // тестовые данные № 1
            {
                name: "zero values",
                values: []byte{},
                want: []interface{}{0, nil},
            },
            // тестовые данные № 2
            {
                name: "mixed values",
                values: []byte("Hello, 世界"),
                want: []interface{}{9, nil},
            },
            {
                name: "error values",
                values: []byte{0xff, 0xfe, 0xfd},
                want: []interface{}{0, ErrInvalidUTF8},
            },
        }
        // перебор всех тестов
        for _, tc := range cases {
            tc := tc
            // запуск отдельного теста
            t.Run(tc.name, func(t *testing.T) {
                // тестируем функцию Sum
                got, err := GetUTFLength(tc.values)
                // проверим полученное значение
                if got != tc.want[0] && err != tc.want[1] {
                        t.Errorf("Sum(%v) = %v, %v; want %v", tc.values, got, err, tc.want)
                }
            })
        }
}