import { Form, useLoaderData } from "react-router";
import type { Route } from "./+types/home";
import { sayHello } from "~/api/hello.server";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "{{project_name}}" },
    { name: "description", content: "{{project_description}}" },
  ];
}

export async function loader({ request }: { request: Request }) {
  try {
    const url = new URL(request.url);
    const name = url.searchParams.get("name");
    if (!name) {
      return { greeting: null };
    }
    const response = await sayHello(name);
    return { greeting: response.message };
  } catch (error) {
    console.error("Error in loader:", error);
    throw error; 
  }
}

export default function Home() {
  const { greeting } = useLoaderData<typeof loader>();
  
  return (
    <div className="min-h-screen flex flex-col">
      <header className="bg-gray-800 text-white p-4 text-center">
        <h1 className="text-2xl">React Router 7 with Go and gRPC Starter</h1>
      </header>
      <main className="flex-1 p-4">
        <div>
          {greeting && <h1 className="text-2xl mb-4">{greeting}</h1>}
          <Form method="get" className="flex items-center">
            <label className="mr-2">Enter your name:</label>
            <input 
              type="text" 
              name="name" 
              placeholder="Enter your name"
              className="border rounded px-2 py-1 mr-2"
            />
            <button 
              type="submit"
              className="px-4 py-1 bg-blue-500 text-white rounded hover:bg-blue-600"
            >
              Say Hello
            </button>
          </Form>
        </div>
      </main>
      <footer className="bg-gray-800 text-white p-4 text-center">
        <p>@sammyjoyce</p>
      </footer>
    </div>
  );
}
