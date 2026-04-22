package main

import (
	"fmt"
)

type Cliente struct {
	ID      int
	Nombre  string
	Carrera string
	Saldo   float64
}

type Producto struct {
	ID        int
	Nombre    string
	Categoria string
	Precio    float64
	Stock     int
}


type Pedido struct {
	ClienteID  int
	ProductoID int
	Cantidad   int
	Total      float64
}



func BuscarClientePorID(clientes []Cliente, id int) int {
	for i, c := range clientes {
		if c.ID == id {
			return i
		}
	}
	return -1
}

func AgregarCliente(clientes []Cliente, nuevo Cliente) []Cliente {
	clientes = append(clientes, nuevo)
	return clientes
}

func EliminarCliente(clientes []Cliente, id int) []Cliente {
	index := BuscarClientePorID(clientes, id)

	if index == -1 {
		return clientes
	}

	clientes = append(clientes[:index], clientes[index+1:]...)
	return clientes
}

func ListarClientes(clientes []Cliente) {
	fmt.Println("\n CLIENTES ")
	for _, c := range clientes {
		fmt.Printf("%d | %s | %s | $%.2f\n",
			c.ID, c.Nombre, c.Carrera, c.Saldo)
	}
}



func BuscarProductoPorID(productos []Producto, id int) int {
	for i, p := range productos {
		if p.ID == id {
			return i
		}
	}
	return -1
}

func AgregarProducto(productos []Producto, nuevo Producto) []Producto {
	productos = append(productos, nuevo)
	return productos
}

func EliminarProducto(productos []Producto, id int) []Producto {
	index := BuscarProductoPorID(productos, id)

	if index == -1 {
		return productos
	}

	productos = append(productos[:index], productos[index+1:]...)
	return productos
}

func ListarProductos(productos []Producto) {
	fmt.Println("\n PRODUCTOS ")
	for _, p := range productos {
		fmt.Printf("%d | %s | %s | $%.2f | Stock: %d\n",
			p.ID, p.Nombre, p.Categoria, p.Precio, p.Stock)
	}
}


func CrearPedido(clientes []Cliente, productos []Producto) {

	var idCliente, idProducto, cantidad int

	fmt.Print("\nIngrese ID Cliente: ")
	fmt.Scan(&idCliente)

	cIndex := BuscarClientePorID(clientes, idCliente)
	if cIndex == -1 {
		fmt.Println("Cliente no existe")
		return
	}

	fmt.Print("Ingrese ID Producto: ")
	fmt.Scan(&idProducto)

	pIndex := BuscarProductoPorID(productos, idProducto)
	if pIndex == -1 {
		fmt.Println("Producto no existe")
		return
	}

	fmt.Print("Cantidad: ")
	fmt.Scan(&cantidad)

	if productos[pIndex].Stock < cantidad {
		fmt.Println("No hay suficiente stock")
		return
	}

	total := float64(cantidad) * productos[pIndex].Precio

	if clientes[cIndex].Saldo < total {
		fmt.Println("Saldo insuficiente")
		return
	}


	productos[pIndex].Stock -= cantidad
	clientes[cIndex].Saldo -= total

	fmt.Printf("Compra realizada. Total: $%.2f\n", total)
}


func main() {

	clientes := []Cliente{
		{ID: 1, Nombre: "Carlos", Carrera: "TI", Saldo: 10},
		{ID: 2, Nombre: "Ana", Carrera: "Civil", Saldo: 15},
		{ID: 3, Nombre: "Harold", Carrera: "Software", Saldo: 20},
	}

	productos := []Producto{
		{ID: 1, Nombre: "Cafe", Categoria: "Bebida", Precio: 1.5, Stock: 10},
		{ID: 2, Nombre: "Cupcake", Categoria: "Dulce", Precio: 3, Stock: 8},
		{ID: 3, Nombre: "Torta", Categoria: "Dulce", Precio: 3, Stock: 10},
		{ID: 4, Nombre: "Capuccino", Categoria: "Bebida", Precio: 3, Stock: 20},
	}

	fmt.Println("Buscar cliente ID 1:", BuscarClientePorID(clientes, 1))
	fmt.Println("Buscar cliente ID 99:", BuscarClientePorID(clientes, 99))

	fmt.Println("\nLista de cliente:")
	ListarClientes(clientes)

	fmt.Println("\nDespués de agregar cliente:")
	nuevoCliente := Cliente{ID: 4, Nombre: "Juan", Carrera: "Medicina", Saldo: 30}
	clientes = AgregarCliente(clientes, nuevoCliente)
	ListarClientes(clientes)

	clientes = EliminarCliente(clientes, 2)
	fmt.Println("\nDespués de eliminar:")
	ListarClientes(clientes)

	fmt.Println("Buscar producto por ID 1:", BuscarProductoPorID(productos, 1))
	fmt.Println("Buscar producto por ID 99:", BuscarProductoPorID(productos, 99))

	fmt.Println("\nLista de Productos")
	ListarProductos(productos)

	fmt.Println("\nDespues de agregar producto")
	nuevoProducto := Producto{ID: 5, Nombre: "Galletas", Categoria: "Dulce", Precio: 2, Stock: 8}
	productos = AgregarProducto(productos, nuevoProducto)
	ListarProductos(productos)

	productos = EliminarProducto(productos, 1)
	fmt.Println("\nProductos después de eliminar:")
	ListarProductos(productos)

	fmt.Println("\n--- CREAR PEDIDO ---")
	CrearPedido(clientes, productos)

	fmt.Println("\nEstado final de clientes:")
	ListarClientes(clientes)

	fmt.Println("\nEstado final de productos:")
	ListarProductos(productos)
}