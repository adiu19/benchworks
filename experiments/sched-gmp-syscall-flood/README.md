# The Goals

1. G >> P, M > P, but only P Ms run Go code at once
    - Lots of goroutines blocked on IO/syscalls.
    - More Ms than Ps (threads created to handle blocking).
    - Still, at any instant, at most GOMAXPROCS threads actually running Go code.
2. P detaches from blocked Ms and is reassigned
    - When a goroutine hits a blocking syscall, M blocks.
    - Runtime detaches P from that M.
    - P is attached to some other M to keep scheduling runnable goroutines.
    - Trace should show Ps hopping between different Ms over time.
3. We visually see the difference between “blocked on syscall” vs “running go”
    - Goroutines spending most time in “syscall” / “sleep”.
    - A smaller subset always runnable.
    - Ps always attached to some runnable M while work exists.

We won’t obsess over exact “10M G / 40 M” numbers because that’s runtime and machine-dependent. The point is to simulate:
- Gs >> Ps
- Ms > Ps
- Only Ps worth of Ms actually run Go code.

# The Approach

1. We flood the system with Gs that quickly block.
2. Each blocking syscall ties up an M temporarily.
3. The runtime is forced to detach Ps and potentially spin up more Ms.
4. A small CPU goroutine continues to make progress on whichever M+P pair is currently scheduled.