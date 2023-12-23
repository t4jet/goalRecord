import { NextResponse } from "next/server";

export async function GET(request) {
  const task = {
    id:1,
    name: "掃除",
  };
  return NextResponse.json(
    { data: task, status: 200 }
  );
}

export async function PUT(request) {
  const body = await request.json();
  return NextResponse.json(
    { status: 200 }
  );
}
