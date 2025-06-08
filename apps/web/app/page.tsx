import Image from 'next/image';

export default function Home() {
  return (
    <div>
      <h1 className="text-2xl font-bold text-center">
        Welcome to the Watch Movie App!
      </h1>
      <hr className="my-4" />
      <Image
        src={'/image.png'}
        alt="Movie Poster"
        width={250}
        height={350}
        className="mx-auto my-4"
      />
    </div>
  );
}
