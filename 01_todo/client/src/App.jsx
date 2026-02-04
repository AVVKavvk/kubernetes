import { useState, useEffect } from "react";

const API_URL = import.meta.env.VITE_BACKEND_API_URL;

function App() {
  const [todos, setTodos] = useState([]);
  const [title, setTitle] = useState("");
  const [date, setDate] = useState("");
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    fetchTodos();
  }, []);

  const fetchTodos = async () => {
    try {
      const response = await fetch(API_URL);
      const data = await response.json();
      setTodos(data || []);
    } catch (error) {
      console.error("Error fetching todos:", error);
    }
  };

  const createTodo = async (e) => {
    e.preventDefault();
    if (!title || !date) return;

    setLoading(true);
    try {
      const response = await fetch(API_URL, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          title,
          date,
          done: false,
        }),
      });

      if (response.ok) {
        setTitle("");
        setDate("");
        fetchTodos();
      }
    } catch (error) {
      console.error("Error creating todo:", error);
    } finally {
      setLoading(false);
    }
  };

  const markDone = async (id) => {
    try {
      await fetch(`${API_URL}/${id}`, {
        method: "PUT",
      });
      fetchTodos();
    } catch (error) {
      console.error("Error marking todo as done:", error);
    }
  };

  const deleteTodo = async (id) => {
    try {
      await fetch(`${API_URL}/${id}`, {
        method: "DELETE",
      });
      fetchTodos();
    } catch (error) {
      console.error("Error deleting todo:", error);
    }
  };

  return (
    <div className="min-h-screen bg-gray-100 py-8 px-4">
      <div className="max-w-2xl mx-auto">
        <h1 className="text-3xl font-bold text-gray-800 mb-8 text-center">
          Todo App
        </h1>

        <form
          onSubmit={createTodo}
          className="bg-white rounded-lg shadow p-6 mb-6"
        >
          <div className="mb-4">
            <label className="block text-gray-700 text-sm font-medium mb-2">
              Title
            </label>
            <input
              type="text"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="Enter todo title"
            />
          </div>

          <div className="mb-4">
            <label className="block text-gray-700 text-sm font-medium mb-2">
              Date
            </label>
            <input
              type="date"
              value={date}
              onChange={(e) => setDate(e.target.value)}
              className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          <button
            type="submit"
            disabled={loading}
            className="w-full bg-blue-500 hover:bg-blue-600 text-white font-medium py-2 px-4 rounded-md transition disabled:bg-blue-300"
          >
            {loading ? "Adding..." : "Add Todo"}
          </button>
        </form>

        <div className="space-y-3">
          {todos.length === 0 ? (
            <p className="text-center text-gray-500 py-8">
              No todos yet. Add one above!
            </p>
          ) : (
            todos.map((todo, index) => (
              <div
                key={index}
                className="bg-white rounded-lg shadow p-4 flex items-center justify-between"
              >
                <div className="flex-1">
                  <h3
                    className={`font-medium ${
                      todo.done ? "line-through text-gray-400" : "text-gray-800"
                    }`}
                  >
                    {todo.title}
                  </h3>
                  <p className="text-sm text-gray-500">{todo.date}</p>
                </div>

                <div className="flex gap-2">
                  {!todo.done && (
                    <button
                      onClick={() => markDone(todo.id)}
                      className="px-3 py-1 bg-green-500 hover:bg-green-600 text-white text-sm rounded transition"
                    >
                      Done
                    </button>
                  )}
                  <button
                    onClick={() => deleteTodo(todo.id)}
                    className="px-3 py-1 bg-red-500 hover:bg-red-600 text-white text-sm rounded transition"
                  >
                    Delete
                  </button>
                </div>
              </div>
            ))
          )}
        </div>
      </div>
    </div>
  );
}

export default App;
