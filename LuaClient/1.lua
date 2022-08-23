--print ("sync src = "..src_dir)
--print ("sync dst = "..dst_dir)

src_dir = arg[1]
dst_dir = arg[2]

sync(src_dir, dst_dir)
