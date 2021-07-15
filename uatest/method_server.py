#!/usr/bin/env python3

from opcua import ua, Server

def even(parent, variant):
    print("even method call with parameters: ", variant.Value)
    ret = (variant.Value % 2 == 0)
    print("even", type(ret))
    return [ua.Variant(ret, ua.VariantType.Boolean)]

def square(parent, variant):
    print("square method call with parameters: ", variant.Value)
    variant.Value *= variant.Value
    return [variant]

def sum_of_squares(parent, variant):
    # print("sum_of_square methos call with parameter: ", variant.Value)
    ret = 3*3 + 8*8
    return [ua.Variant(ret, ua.VariantType.Int64)]

if __name__ == "__main__":
    server = Server()
    server.create_custom_data_type(2, "ComplexType", description="A complex type")

    server.set_endpoint("opc.tcp://0.0.0.0:4840/")

    ns = server.register_namespace("http://gopcua.com/")
    main = server.nodes.objects.add_object(ua.NodeId("main", ns), "main")
    fnEven = main.add_method(ua.NodeId("even", ns), "even", even, [ua.VariantType.Int64], [ua.VariantType.Boolean])
    fnSquare = main.add_method(ua.NodeId("square", ns), "square", square, [ua.VariantType.Int64], [ua.VariantType.Int64])
    fnSumOfSquare = main.add_method(ua.NodeId("sumOfSquare", ns), "sumOfSquare", sum_of_squares, [ua.VariantType.ExtensionObject], [ua.VariantType.Int64])

    server.start()
