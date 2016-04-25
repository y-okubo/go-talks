# fiddle というライブラリでも OK か？
# http://docs.ruby-lang.org/ja/2.1.0/library/fiddle.html

require "ffi"

module Fib
  extend FFI::Library
  ffi_lib "libgofib.so"
  attach_function :fib, [:uint], :uint
end

puts Fib.fib(40)
