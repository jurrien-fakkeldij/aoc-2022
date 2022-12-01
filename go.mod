module main

go 1.19

require internal/filereader v1.0.0

replace internal/filereader => ./internal/filereader

require internal/transformer v1.0.0

replace internal/transformer => ./internal/transformer
