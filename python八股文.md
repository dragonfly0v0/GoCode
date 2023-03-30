python面试

# 1.读取文件sys.path

# 2.协程和线程

Python中的协程和线程都是实现多任务的方式，但它们的运行机制和特点有所不同。

线程是操作系统中最小的执行单元，线程之间是抢占式调度的，即一个线程执行完毕后，操作系统会立即切换到另一个线程上去执行。线程之间共享进程的资源，如内存空间、文件句柄等，因此需要进行锁定和同步来避免资源竞争和数据损坏。多线程编程需要考虑锁的使用、线程间通信等问题，编写起来比较复杂。

协程是一种轻量级的线程，也称为微线程，是由程序自身控制的，与操作系统无关。协程在同一个线程内，多个协程之间共享一个线程，它们之间的切换是由程序自身控制的，可以避免线程切换的开销，因此协程的执行效率比线程高。协程的切换是通过yield、await等关键字实现的，通过在协程执行过程中主动释放控制权，让出CPU资源，然后转而执行其他协程，从而实现多任务并发执行。

在Python 3.5及以上版本中，引入了async/await语法，使得协程编程更加方便。通过async/await语法，可以简单地定义协程函数，使用异步的方式进行I/O操作，从而实现高效的异步编程。

总的来说，协程比线程更加轻量级、高效、灵活，适合处理I/O密集型的任务，如网络通信、数据库访问等。而线程则适合处理CPU密集型的任务，如计算、数据处理等。

# 3.反射

在 Python 中，反射指的是通过字符串来访问、操作对象的属性或方法。这种通过字符串来实现对象访问、操作的能力，在 Python 中非常重要，它让 Python 成为一门非常灵活和动态的语言。

Python 提供了几个内置函数来实现反射的功能，包括：

- `hasattr(object, name)`：判断一个对象是否有指定的属性或方法。
- `getattr(object, name[, default])`：获取一个对象的属性或方法，如果没有则返回指定的默认值。
- `setattr(object, name, value)`：设置一个对象的属性或方法。
- `delattr(object, name)`：删除一个对象的属性或方法。

除了上述内置函数之外，Python 中还有一些特殊的魔术方法（Magic Method）也可以用来实现反射的功能。比如：

- `__getattr__(self, name)`：当获取一个不存在的属性时会调用该方法。
- `__setattr__(self, name, value)`：当设置一个属性时会调用该方法。
- `__delattr__(self, name)`：当删除一个属性时会调用该方法。

下面是一个示例代码，展示了如何使用反射来访问、操作对象的属性和方法：

```python
class Person:
    def __init__(self, name, age):
        self.name = name
        self.age = age
    
    def say_hello(self):
        print(f"{self.name} says hello!")

p = Person("Alice", 18)

# 使用 hasattr 判断对象是否有指定属性或方法
if hasattr(p, "name"):
    print(getattr(p, "name")) # 使用 getattr 获取对象属性
if hasattr(p, "say_hello"):
    getattr(p, "say_hello")() # 使用 getattr 调用对象方法

# 使用 setattr 设置对象属性
setattr(p, "gender", "female")
print(p.gender)

# 使用 delattr 删除对象属性
delattr(p, "age")
if not hasattr(p, "age"):
    print("Age has been deleted from person object.")

```

# 4.元类

在 Python 中，元类（Metaclass）是一种特殊的类，用于创建其他类的类。换句话说，元类是用来控制类创建的过程的。

在 Python 中，类的创建过程分为两个阶段：

1. 解析阶段（Compile-Time）：解析类的代码，将类定义转换成内部表示形式（即字节码）。
2. 运行阶段（Runtime）：根据类的定义创建类对象（Class Object），并将该对象绑定到类名上。

元类的作用是可以控制这两个阶段的行为，从而实现一些高级的功能。比如，我们可以通过元类来自动注册子类，实现类似 ORM 框架的功能。

在 Python 中，可以使用 `type()` 内置函数来定义元类，也可以通过继承 `type` 类来自定义元类。当我们使用自定义元类创建一个新类时，Python 会调用元类的 `__new__` 和 `__init__` 方法来完成类的创建过程。其中，`__new__` 方法负责创建类对象，而 `__init__` 方法负责初始化类对象。

下面是一个简单的示例代码，演示了如何定义一个元类并使用它来创建一个新类：

```python
class MyMeta(type):
    def __new__(cls, name, bases, attrs):
        # 自定义处理逻辑
        print("Creating class...")
        attrs["greeting"] = "Hello, world!"
        return super().__new__(cls, name, bases, attrs)
    
class MyClass(metaclass=MyMeta):
    pass

print(MyClass.greeting) # 输出 "Hello, world!"

```

在这个示例中，我们定义了一个名为 `MyMeta` 的元类，并在其中重载了 `__new__` 方法。在 `__new__` 方法中，我们添加了一个属性 `greeting` 到类的属性字典中，并返回了新的类对象。

接下来，我们使用 `MyClass` 类，并将 `MyMeta` 元类作为它的元类来创建一个新类。在创建过程中，Python 会自动调用 `MyMeta` 的 `__new__` 和 `__init__` 方法。最终，我们通过访问 `MyClass.greeting` 属性来验证新类是否被成功地创建。

# 4.元组解包

在 Python 中，元组解包（Tuple unpacking）是一种将元组中的多个值同时赋值给多个变量的方式。使用元组解包可以使代码更加简洁明了，同时也更加高效。

```python
# 定义一个包含多个值的元组
t = (1, 2, 3)

# 使用元组解包将元组中的多个值赋值给多个变量
a, b, c = t

# 打印变量的值
print(a) # 输出 1
print(b) # 输出 2
print(c) # 输出 3

```

需要注意的是，如果元组中的值的数量与变量的数量不匹配，会抛出 `ValueError` 异常。另外，我们还可以使用 `*` 运算符来进行扩展的元组解包，如下面的示例代码所示：

```python
# 定义一个包含多个值的元组
t = (1, 2, 3, 4)

# 使用扩展的元组解包将前两个值赋值给两个变量，剩余的值赋值给一个变量
a, b, *c = t

# 打印变量的值
print(a) # 输出 1
print(b) # 输出 2
print(c) # 输出 [3, 4]

```

在这个示例中，我们使用扩展的元组解包将前两个元素赋值给两个变量 `a` 和 `b`，而将剩余的值赋值给变量 `c`。在打印 `c` 变量的值时，我们得到了一个列表 `[3, 4]`，其中包含了元组中剩余的两个值。

# 内存池

在 Python 中，使用内存池机制可以通过 `sys.getsizeof()` 函数来观察。下面是一个示例代码，演示了如何使用 `sys.getsizeof()` 函数观察内存池的情况：

```python
import sys

# 定义一个空列表
a = []

# 打印列表占用的内存大小
print(sys.getsizeof(a))  # 输出 64

# 向列表中添加 10 个元素
for i in range(10):
    a.append(i)

# 再次打印列表占用的内存大小
print(sys.getsizeof(a))  # 输出 192

# 从列表中删除 5 个元素
for i in range(5):
    a.pop()

# 再次打印列表占用的内存大小
print(sys.getsizeof(a))  # 输出 192

# 清空列表
a.clear()

# 再次打印列表占用的内存大小
print(sys.getsizeof(a))  # 输出 64

```

# 内存优化

Python 内存优化通常有以下几个方面：

1. 对象池：Python 内部会对一些常用的对象，例如整数、字符串、布尔值等进行缓存和重复利用，避免频繁创建和销毁。可以使用 sys.intern() 对字符串进行缓存，或者使用 from sys import intern 进行全局缓存。

2. 避免循环引用：在 Python 中，当一个对象的引用计数为 0 时，它会被 Python 的垃圾回收机制回收。而循环引用会导致引用计数永远不为 0，导致内存泄漏。可以使用 weakref 模块或者手动将循环引用断开来避免这种情况。

3. 减少拷贝：Python 的内存分配和释放都比较昂贵，因此减少不必要的对象拷贝可以提高程序的效率。可以使用切片、可变对象的 clear() 方法、map() 等方法来减少拷贝。

4. 使用生成器：生成器是一种延迟计算的方式，可以在需要时才生成计算结果，避免一次性生成大量对象。

5. 使用内置函数和标准库：Python 内置函数和标准库经过优化，使用它们可以提高程序的效率，避免重复造轮子。例如使用列表推导式、map()、filter() 等高效的内置函数。

6. 使用 C 扩展或者 Cython：C 扩展或者 Cython 可以提高 Python 程序的性能，避免一些 Python 解释器的开销，但是需要编写 C 或者 Cython 代码，对于一些小项目或者脚本并不一定值得使用。

总之，Python 内存优化的原则是尽量避免不必要的内存分配和释放，减少对象拷贝，利用好 Python 内置的优化和标准库，对于一些性能要求高的场景，可以考虑使用 C 扩展或者 Cython。

# GIL  全局解释锁



4. 闭包
   
   Python 闭包是一种函数嵌套的结构，在一个函数内部定义另一个函数，并且返回该函数，从而形成一个闭合的空间，内部函数可以访问外部函数的变量和参数。这种结构在函数式编程中非常常见，可以用来实现许多高级功能。
   
   具体来说，当一个函数内部定义了另一个函数，并且外部函数将该内部函数作为返回值时，内部函数就成为了闭包。在闭包内部，可以访问外部函数的变量和参数，但是这些变量和参数会被保存在闭包中，即使外部函数已经执行完毕，闭包仍然可以访问这些变量和参数。这种特性可以用来实现一些高级功能，例如装饰器、工厂函数、缓存等。
   
   以下是一个简单的闭包示例：
   
   ```python
   def outer_func(x):
       def inner_func(y):
           return x + y
       return inner_func
   
   closure = outer_func(10)
   print(closure(20))  # 输出 30
   
   ```
   
   在上面的例子中，outer_func 是一个外部函数，它接受一个参数 x，然后在内部定义了一个函数 inner_func，inner_func 接受一个参数 y，返回 x + y 的值。最后，outer_func 返回 inner_func 函数对象。
   
   在调用 outer_func(10) 之后，我们得到一个闭包对象 closure，它实际上就是 inner_func 函数对象。当我们调用 closure(20) 时，它会返回 30，这是因为 x 在外部函数中被设置为 10，然后被保存在闭包中，闭包可以访问并使用它。因此，这个闭包实现了一个简单的加法器，它可以记住一个值 x，然后用于后续的加法运算。

5. 上下文管理器
   
   在 Python 中，上下文管理器（Context Manager）是一种处理资源的机制，它可以在代码块执行前获取资源，并在代码块执行完毕后释放资源。它通过 with 语句来实现，with 语句可以确保在代码块执行完毕后自动释放资源，无论代码块是否正常执行，都可以保证资源的释放。
   
   一个对象如果实现了 **enter**() 和 **exit**() 两个方法，那么这个对象就是一个上下文管理器。**enter**() 方法在代码块执行前被调用，它可以获取资源并返回一个值，这个值会被赋给 as 关键字后面的变量。**exit**() 方法在代码块执行完毕后被调用，它可以释放资源，并且可以处理可能出现的异常。
   
   以下是一个简单的上下文管理器示例：
   
   ```python
   class MyContextManager:
       def __enter__(self):
           print('Enter')
           return 'Hello'
   
       def __exit__(self, exc_type, exc_value, traceback):
           print('Exit')
   
   with MyContextManager() as cm:
       print(cm)
   ```
   
   在这个示例中，我们定义了一个 MyContextManager 类，它实现了 **enter**() 和 **exit**() 两个方法。在 with 语句中，我们创建了一个 MyContextManager 对象，然后调用 **enter**() 方法，获取到了一个字符串 'Hello'，并且将它赋给了 cm 变量。接着，我们打印了 cm 变量，然后代码块执行完毕，自动调用了 **exit**() 方法，释放了资源，并打印了 'Exit'。
   
   上下文管理器可以用来管理各种资源，例如文件、锁、网络连接等。使用上下文管理器可以确保资源的正确获取和释放，同时也可以让代码更加简洁和易读。
