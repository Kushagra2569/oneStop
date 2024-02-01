<script>
  import { NewTodo } from "../../wailsjs/go/main/Todos.js";
  import { GetTodos } from "../../wailsjs/go/main/Todos.js";

  let val = "Hello World";
  function addTodo() {
    NewTodo("Sample todo", 1).then((res) => {
      val = res;
    });
  }

  function print() {
    GetTodos().then((res) => {
      val = res;
    });
  }
</script>

<main>
  <p>{val}</p>
  <button on:click={addTodo}>Add Todo</button>
  <button on:click={print}>print Todo</button>
</main>
