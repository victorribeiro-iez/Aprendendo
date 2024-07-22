package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Cliente struct {
	id        int64
	nome      string
	sobrenome string
}

func OpenConn() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=Cliente sslmode=disable")

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	return db, err
}

func Get(id int) (c Cliente, err error) {

	conn, err := OpenConn()

	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`select id,nome,sobrenome from public.tb_cliente where id=$1`, id)

	err = row.Scan(&c.id, &c.nome, &c.sobrenome)

	if err != nil {
		return c, err
	}

	return c, nil

}

func Delete(id int) (int64, error) {
	conn, err := OpenConn()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM public.tb_cliente WHERE id = $1`, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func Insert(c Cliente) (id int64, err error) {
	conn, err := OpenConn()
	if err != nil {
		return
	}
	defer conn.Close()
	sql := `INSERT INTO tb_cliente (id, nome, sobrenome) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, c.id, c.nome, c.sobrenome).Scan(&c.id)

	if err != nil {
		return
	}

	return id, nil

}

func Update(id int64, c Cliente) (int64, error) {
	conn, err := OpenConn()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	sql := `UPDATE public.tb_cliente SET nome = $2, sobrenome = $3 WHERE id = $1`
	res, err := conn.Exec(sql, id, c.nome, c.sobrenome)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func main() {

	/*Get id*/
	/*
		usuario, erro := Get(2)
		fmt.Println(usuario)
		if erro != nil {
			fmt.Println("Error:", erro)
			return
		}*/

	/*Post*/
	/*
		cliente := Cliente{
			id:        3,
			nome:      "Teste nome",
			sobrenome: "Teste sobrenome",
		}

		id, erro := Insert(cliente)
		fmt.Println("ID do novo cliente:", id)


		if erro != nil {
			fmt.Println("Error:", erro)
			return
		}

	*/

	/*Delete*/
	/*
		affectedRows, err := Delete(1)
		if err != nil {
			fmt.Println("Erro ao deletar:", err)
			return
		}
		fmt.Println("Número de linhas afetadas:", affectedRows)
	*/

	/*Update*/
	/*
		cliente := Cliente{
			nome:      "Victor Alexandre Teste",
			sobrenome: "Teste sobrenome",
		}

		affectedRows, err := Update(1, cliente)
		if err != nil {
			fmt.Println("Erro ao atualizar:", err)
			return
		}

		fmt.Println("Número de linhas afetadas:", affectedRows)
	*/

}
