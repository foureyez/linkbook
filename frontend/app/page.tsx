import Image from "next/image";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div>
        <form>
          <input type="url" name="url" />
          <button type="submit">Add bookmark</button>
        </form>
      </div>
    </main>
  );
}
