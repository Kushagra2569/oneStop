<script>
  import { NewTodo } from "../../wailsjs/go/main/App.js";
  import { GetTodo } from "../../wailsjs/go/main/Todo.js";

  let val = "Hello World";
  function addTodo() {
    NewTodo("Sample todo", 1).then((res) => {
      val = res;
    });
  }

  function print() {
    GetTodo().then((res) => {
      val = res;
    });
  }
</script>

<main>
  <button on:click={addTodo}>Add Todo</button>
  <p>{val}</p>
  <button on:click={print}>print Todo</button>
</main>
