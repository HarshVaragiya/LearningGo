# Struct
- Go does not have classes
- See [Methods.go](Methods.go) to see how to define methods for a given struct

- You can declare a method on non-struct types, too.
- In [Methods2.go](Methods2.go) we see a numeric type MyFloat with an Abs method
- You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).

## Pointer receiver

- There are two reasons to use a pointer receiver.
    - The first is so that the method can modify the value that its receiver points to.

    - The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

In [PointerReceiver.go](PointerReceiver.go) both Scale and Abs are with receiver type *Vertex, even though the Abs method needn't modify its receiver.

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.