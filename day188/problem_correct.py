def make_functions():
    flist = []

    for i in [1, 2, 3]:
        def build_print_i(i):
            def print_i():
                print(i)
            return print_i
        flist.append(build_print_i(i))

    return flist

functions = make_functions()
for f in functions:
    f()
