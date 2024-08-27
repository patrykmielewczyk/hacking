import os

func_name = "system"
command = getattr(os, func_name)
command("id")
